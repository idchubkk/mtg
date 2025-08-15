package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(Lang[CurrentLang]["usage"])
		return
	}
	cmd := os.Args[1]
	cfg, _ := LoadConfig("config.json")

	switch cmd {
	case "install":
		InstallMTProxy()
	case "start":
		StartProxy(cfg)
	case "stop":
		StopProxy()
	case "restart":
		RestartProxy(cfg)
	case "setport":
		if len(os.Args) < 3 {
			fmt.Println("请指定端口")
			return
		}
		port, _ := strconv.Atoi(os.Args[2])
		SetPort(cfg, port)
	case "setkey":
		if len(os.Args) < 3 {
			fmt.Println("请指定密钥")
			return
		}
		SetSecret(cfg, os.Args[2])
	default:
		fmt.Println(Lang[CurrentLang]["usage"])
	}
}
