package config

import (
	"encoding/json"
	"log"
	"os"
)

/* Change default values in default.go! */
type Config struct {
	SQLite   SQLite
	MySQL    MySQL
	DbDriver string // use the above SQL or MySQL configuration? valid options are "mysql" or "sqlite"
	Port     int    // which port to host the API REST server on

}

// If the file is not found, (nil, nil) is returned to make it easy for the
// caller to decide to create one or whatever.
func Load(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	defer f.Close()

	conf := new(Config)
	err = json.NewDecoder(f).Decode(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func Log(thing interface{}) {
	if output, err := json.MarshalIndent(thing, "", "\t"); err == nil {
		log.Printf("%s\n", output)
	} else {
		log.Printf("%s\n", err.Error())
	}
}
