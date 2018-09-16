package config

import (
	"encoding/json"
	"os"
)

func DefaultConfig() *Config {
	conf := &Config{
		SQLite: SQLite{
			UseMemory: false,
			FilePath:  "tmn-api.sqlite",
		},
		MySQL: MySQL{
			Host:     "localhost",
			Port:     3306,
			Username: "DevUser",
			Password: "DevPassword",
			Database: "development",
		},
		DbDriver: "mysql",
		Port:     8080,
	}
	return conf
}

func WriteDefault(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(DefaultConfig())
	if err != nil {
		return err
	}

	return nil
}
