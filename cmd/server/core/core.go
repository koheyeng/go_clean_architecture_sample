package core

import (
	"fmt"

	"github.com/koheyeng/go_clean_architecture_sample/driver/rdb"
	"github.com/koheyeng/go_clean_architecture_sample/logger"
)

type Core struct {
	APIBaseURL     string
	APIInsecureTLS bool

	DBHandler   rdb.DBHandler
	BindAddress string

	Logger logger.Logger
}

func (c *Core) Valid() error {
	switch {
	case c == nil:
		return fmt.Errorf("`Config` not present")
	case c.DBHandler.DB == nil:
		return fmt.Errorf("`DB` not present")
	case c.BindAddress == "":
		return fmt.Errorf("`BindAddress` not present")
	}

	return nil
}
