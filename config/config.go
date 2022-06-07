package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConfigStruct struct {
	Token     string   `json:"Token"`
	Message   string   `json:"Message"`
	BlackList []string `json:"BlackList"`
}

func LoadConfig() *ConfigStruct {
	log.Print("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var loadedConfig = ConfigStruct{}
	err = json.Unmarshal(file, &loadedConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &loadedConfig
}
