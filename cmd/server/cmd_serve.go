package server

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/koheyeng/go_clean_architecture_sample/cmd"
	"github.com/koheyeng/go_clean_architecture_sample/cmd/config"
)

type CmdServe struct {
	cmd.Meta
}

func (c *CmdServe) Synopsis() string {
	return "Surve Sample App"
}

func (c *CmdServe) Help() string {
	txt := fmt.Sprintf(`
Usage:
	%s %s

Description:
	%s
`,
		c.CmdName, c.SubCmdName, c.Synopsis())

	return txt[1:]
}

func (c *CmdServe) Run(args []string) int {
	configFilePath := config.ConfigFilePath
	// logger := logger.New("")

	flagSet := flag.NewFlagSet(c.CmdName+" "+c.SubCmdName, flag.ExitOnError)
	flagSet.Usage = func() { fmt.Fprintf(c.Meta.UI.ErrorWriter, "%s\n", c.Help()) }
	flagSet.StringVar(&configFilePath, "c", configFilePath, "")
	flagSet.StringVar(&configFilePath, "config", configFilePath, "")
	flagSet.Parse(args)

	conf := config.NewConfig()

	configFile, err := os.Open(configFilePath)
	if err != nil {
		if configFilePath != config.ConfigFilePath || err.(*os.PathError).Err != syscall.ENOENT {
			c.Meta.Errorf("%v\n", err)
			return 1
		}
	} else {
		if err := conf.ReadConfig(configFile); err != nil {
			c.Meta.Errorf("%v\n", err)
			return 1
		}
	}

	// dbHandler, err := rdb.NewDatabase(conf.DB)

	return 0
}
