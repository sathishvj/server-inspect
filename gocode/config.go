package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const configFilePath = "config.json"

type config struct {
	Port  int
	Files []InspectFileCfg
}

type InspectFileCfg struct {
	Dir, Pattern string
	Recursive    bool
}

func NewConfig(filepath string) (*config, error) {
	//open the json file and put in entries
	b, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("config.go:NewConfig(): Error reading file: %v\n", e)
		return nil, e
	}

	var cfg config
	err := json.Unmarshal(b, &cfg)
	fmt.Printf("config.go:NewConfig(): Read config data: %+v\n", cfg)
	return &cfg, err
}
