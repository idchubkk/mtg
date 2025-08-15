package config

import (
	"os"
	"strings"
)

type Config struct {
	Secret     string
	BindTo     string
	Blocklist  []string
	ChainProxy string
}

func LoadConfig(path string) (*Config, error) {
	// 简化版，不用第三方库，直接读取文本
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "secret":
			cfg.Secret = value
		case "bind_to":
			cfg.BindTo = value
		case "blocklist":
			cfg.Blocklist = strings.Split(value, ",")
		case "chain_proxy":
			cfg.ChainProxy = value
		}
	}

	return cfg, nil
}
