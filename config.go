package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port   int    `json:"port"`
	Secret string `json:"secret"`
	Path   string `json:"path"` // MTProxy 可执行文件路径
}

func LoadConfig(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return &Config{
			Port:   443,
			Secret: "",
			Path:   "./mtproto-proxy",
		}, nil
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return &cfg, err
}

func SaveConfig(file string, cfg *Config) error {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(file, data, 0644)
}
