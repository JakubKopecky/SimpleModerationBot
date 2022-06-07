package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigStruct struct {
	Token     string
	Message   string   `json:"Message"`
	BlackList []string `json:"BlackList"`
}

func LoadConfig() *ConfigStruct {
	log.Print("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("!!!")

	var loadedConfig = ConfigStruct{}
	err = json.Unmarshal(file, &loadedConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print(loadedConfig)

	loadedConfig.Token = os.Getenv("BOT_TOKEN")
	log.Print("Token is " + loadedConfig.Token)

	log.Print(loadedConfig)

	return &loadedConfig
}
