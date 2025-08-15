package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port   int    `json:"port"`
	Secret string `json:"secret"`
	Path   string `json:"path"`
}

var cfgFile = "config.json"

func LoadConfig() *Config {
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		return &Config{Port: 443, Secret: "", Path: "./MTProxy/objs/bin/mtproto-proxy"}
	}
	var cfg Config
	json.Unmarshal(data, &cfg)
	return &cfg
}

func SaveConfig(cfg *Config) {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(cfgFile, data, 0644)
}
