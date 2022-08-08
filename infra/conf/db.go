package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfDB struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	DbName  string `json:"db-name"`
	Charset string `json:"charset"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
}

func ReadConfDB() (*ConfDB, error) {
	const conffile = "infra/conf/db.json"
	conf := new(ConfDB)
	readFile, err := ioutil.ReadFile(conffile)
	if err != nil {
		return conf, fmt.Errorf("failed to read json conf file: %w", err)
	}
	err = json.Unmarshal([]byte(readFile), conf)
	if err != nil {
		return conf, fmt.Errorf("failed to unmarshal conf file: %w", err)
	}
	return conf, nil
}
