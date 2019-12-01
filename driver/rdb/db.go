package rdb

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/koheyeng/go_clean_architecture_sample/cmd/config"
)

type DBHandler struct {
	DB *gorm.DB
}

func NewDatabase(c *config.DBConfig) (*DBHandler, error) {

	conn := fmt.Sprintf("host=%s port=%s sslmode=%s dbname=%s user=%s password=%s", c.Host, c.Port, c.SSLMode, c.DBName, c.User, c.Password)
	db, err := gorm.Open(c.Driver, conn)
	if err != nil {
		return nil, xerrors.Errorf("DB connection open error: %w", err)
	}

	db.SingularTable(true)
	db.LogMode(c.LogMode)

	dbHandler := new(DBHandler)
	dbHandler.DB = db

	return dbHandler, nil
}
