package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port        uint16
	Source      string
	Destination string
}

func NewConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	configuration := Config{}
	err = json.Unmarshal(file, &configuration)

	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
