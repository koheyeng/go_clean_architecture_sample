package config

import (
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var ConfigFilePath = filepath.Base(os.Args[0]) + ".conf"

type Config struct {
	API *APIConfig
	DB  *DBConfig
}

type APIConfig struct {
	BaseURL     string `toml:"baseurl"`
	InsecureTLS bool   `toml:"insecure-tls"`
	BindAddress string `toml:"bind-address"`
}

type DBConfig struct {
	Driver   string `toml:"driver"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	SSLMode  string `toml:"sslmode"`
	DBName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"passowrd"`
	LogMode  bool   `toml:"logmode"`
}

func NewConfig() *Config {
	c := &Config{
		API: &APIConfig{
			BaseURL:     "http://127.0.0.1:8080",
			InsecureTLS: false,
			BindAddress: "127.0.0.1:3000",
		},
		DB: &DBConfig{
			Driver:   "postgres",
			Host:     "127.0.0.1",
			Port:     "5432",
			SSLMode:  "disable",
			DBName:   "sample",
			User:     "sample",
			Password: "p@ssw0rd",
		},
	}

	return c
}

func (c *Config) ReadConfig(rd io.Reader) error {
	if _, err := toml.DecodeReader(rd, c); err != nil {
		return err
	}

	return nil
}
