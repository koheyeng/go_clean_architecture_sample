package cmd

import (
	"fmt"

	"github.com/mitchellh/cli"
)

// Meta struct
type Meta struct {
	UI         cli.BasicUi
	CmdName    string
	SubCmdName string
	Version    string
}

// Errorf prints the error to ErrorWriter with the prefix string
func (m Meta) Errorf(format string, v ...interface{}) {
	fmt.Fprintf(m.UI.ErrorWriter, "[ERR]"+format, v...)
}
