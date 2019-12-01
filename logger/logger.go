package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Logger struct {
	module string
	logger *log.Logger
}

func New(prefix string) *Logger {
	return &Logger{
		logger: log.New(os.Stdout, prefix, log.Ldate|log.Lmicroseconds),
	}
}

func (l *Logger) Print(v interface{}) {
	l.logger.Output(2, fmt.Sprintf("{\"%s\": {\"log\": {\"message\": \"%s\"}}}", l.module, v))
}

func (l *Logger) PrintJson(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		l.PrintError(0, err)
	}
	l.logger.Output(2, fmt.Sprintf("{\"%s\": {\"log\": {\"message\": \"%s\"}}}", l.module, string(b)))
}

func (l *Logger) SetModuleName(module string) {
	l.module = module
}

func (l *Logger) PrintFatal(code int, err error) {
	proxyError := NewSampleError(Fatal, code, fmt.Sprintf("modules=%s\n", l.module))
	proxyError.Wrap(err).Print()
}

func (l *Logger) PrintError(code int, err error) {
	proxyError := NewSampleError(Error, code, fmt.Sprintf("modules=%s\n", l.module))
	proxyError.Wrap(err).Print()
}
func (l *Logger) PrintWarning(code int, err error) {
	proxyError := NewSampleError(Warning, code, fmt.Sprintf("modules=%s\n", l.module))
	proxyError.Wrap(err).Print()
}
