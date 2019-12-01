package logger

import (
	"fmt"
	"log"

	"golang.org/x/xerrors"
)

type ErrLevel int

//go:generate stringer -type=ErrLevel
const (
	Fatal ErrLevel = iota
	Error
	Warning
)

type SampleError struct {
	level ErrLevel
	code  int
	msg   string
	err   error
	frame xerrors.Frame
}

func NewSampleError(level ErrLevel, code int, msg string) *SampleError {
	return &SampleError{
		level: level,
		code:  code,
		msg:   msg,
		err:   nil,
	}
}

func (e SampleError) Wrap(next error) *SampleError {
	e.err = next
	return &e
}

func (e *SampleError) Error() string {
	return fmt.Sprintf("%s: code=%d, %s", e.level, e.code, e.msg)
}

func (e *SampleError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *SampleError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.Error())
	e.frame.Format(p)
	return e.err
}

func (e *SampleError) Is(err error) bool {
	var appErr *SampleError
	return xerrors.As(err, &appErr) && e.code == appErr.code
}

func (e *SampleError) Print() {
	switch e.level {
	case Fatal:
		log.Fatal("%+v", e)
	case Error:
		log.Printf("%+v", e)
	case Warning:
		log.Printf("%+v", e)
	}
}
