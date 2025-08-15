package main

import (
	"flag"
)

type Config struct {
	Port int
}

func ParseConfig() *Config {
	port := flag.Int("port", 4433, "监听端口")
	flag.Parse()
	return &Config{
		Port: *port,
	}
}
