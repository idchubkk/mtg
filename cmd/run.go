package main

import (
	"fmt"
	"log"
	"my-tg-proxy/config"
	"my-tg-proxy/proxy"
)

func run() {
	cfg, err := config.LoadConfig("example.toml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	fmt.Println("Starting MTProto Proxy on", cfg.BindTo)
	p := proxy.NewProxy(cfg)
	if err := p.Start(); err != nil {
		log.Fatal("Proxy error:", err)
	}
}
