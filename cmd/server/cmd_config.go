package server

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/koheyeng/go_clean_architecture_sample/cmd"
	"github.com/koheyeng/go_clean_architecture_sample/cmd/config"
)

type CmdConfig struct {
	cmd.Meta
}

func (c *CmdConfig) Synopsis() string {
	return "Write a default configuration to stdout."
}

func (c *CmdConfig) Help() string {
	txt := fmt.Sprintf(`
Usage:
	%s %s

Description:
	%s
`,
		c.CmdName, c.SubCmdName, c.Synopsis())

	return txt[1:]
}

func (c *CmdConfig) Run(args []string) int {
	config := config.NewConfig()

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		c.Errorf("%v\n", err)
		return 1
	}

	fmt.Fprintf(c.UI.Writer, "%s", buf.String())
	return 0
}
