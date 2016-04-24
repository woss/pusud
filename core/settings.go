package core

import (
	"os"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var settingsFilename = "settings.yaml"

type Settings struct {
	Authenticator string `yaml:"authenticator"`
}

func ReadSettings() *Settings {
	s := Settings{}

	data, err := ioutil.ReadFile(settingsFilename)

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Settings file %s not found. Try copying settings.example.yaml", settingsFilename)
		} else {
			log.Fatalf("Error reading settings: %v", err)
		}
	}

	log.Printf("Read settings from %s", settingsFilename)

	yaml.Unmarshal(data, &s)

	return &s
}
