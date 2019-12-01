package server

import (
	"fmt"

	"github.com/koheyeng/go_clean_architecture_sample/cmd"
)

type CmdVersion struct {
	cmd.Meta
}

func (c *CmdVersion) Synopsis() string {
	return "Display version info."
}

func (c *CmdVersion) Help() string {
	txt := fmt.Sprintf(`
Usage:
	%s %s

Description:
	%s
`,
		c.CmdName, c.SubCmdName, c.Synopsis())

	return txt[1:]
}

func (c *CmdVersion) Run(args []string) int {
	fmt.Println(c.Version)
	return 0
}
