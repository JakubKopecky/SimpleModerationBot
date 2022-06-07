package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	Token     string
	BlackList *[]string

	config *configStruct
)

type configStruct struct {
	Token     string   `json:"Token"`
	BlackList []string `json:"BlackList"`
}

func ReadConfig() error {
	log.Print("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Panic(err)
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Panic(err)
		return err
	}

	Token = config.Token
	BlackList = &config.BlackList

	return nil
}