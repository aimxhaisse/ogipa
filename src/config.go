package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ListenOn string // Interface to listen on
	Rethink  ConfigRethink
}

type ConfigRethink struct {
	Address  string
	Database string
}

// NewConfig returns a configuration from file pointer by path
func NewConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
