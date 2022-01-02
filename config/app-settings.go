package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var app_settings map[string]string

func InitSettings() {
	file, err := ioutil.ReadFile("config/seed.json")

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(file, &app_settings)

	log.Println("Initialized settings")
}

func GetSetting(i string) string {
	return app_settings[i]
}
