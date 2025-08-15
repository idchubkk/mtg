package main

import (
	"fmt"
)

func main() {
	cfg := ParseConfig()
	secretStr := GenerateSecret()
	secret := []byte(secretStr)

	fmt.Println("生成的 MTProxy secret:", secretStr)
	fmt.Printf("代理链接示例: tg://proxy?server=YOUR_SERVER_IP&port=%d&secret=%s\n", cfg.Port, secretStr)

	StartProxy(cfg.Port, secret)
}
