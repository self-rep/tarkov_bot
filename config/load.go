package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	AmmoData    *AmmoStruct
	BotSettings *Settings
)

func LoadFile() error {
	f, err := ioutil.ReadFile("./assets/ammo.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(f, &AmmoData)
	if err != nil {
		return err
	}
	return nil
}

func LoadSettings() error {
	f, err := ioutil.ReadFile("./assets/settings.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(f, &BotSettings)
	if err != nil {
		return err
	}
	return nil
}
