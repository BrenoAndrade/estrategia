package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// ConfigFile struct for config file
type ConfigFile struct {
	Port             string
	ConnectionString string
}

// GetConfig get config file
func GetConfig() ConfigFile {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("[ERROR] read config file.", err)
		os.Exit(0)
	}

	config := ConfigFile{}
	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		log.Println("[ERROR] create struct config file.", err)
		os.Exit(0)
	}

	return config
}
