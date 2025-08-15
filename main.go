package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 语言选择
	fmt.Println("Select language / 言語を選択 / 选择语言 / Selecione a língua:")
	fmt.Println("1: English")
	fmt.Println("2: 中文")
	fmt.Println("3: 日本語")
	fmt.Println("4: língua portuguesa")
	fmt.Print("Enter number [default 1]: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		input = "1"
	}

	switch input {
	case "1":
		CurrentLang = "en"
	case "2":
		CurrentLang = "zh"
	case "3":
		CurrentLang = "jp"
	case "4":
		CurrentLang = "pt"
	default:
		fmt.Println("Invalid input, defaulting to English")
		CurrentLang = "en"
	}

	fmt.Printf("Current language: %s\n\n", CurrentLang)

	// 显示可选操作
	fmt.Println(Lang[CurrentLang]["usage"])
	for {
		fmt.Print("Command: ")
		cmdLine, _ := reader.ReadString('\n')
		cmdLine = strings.TrimSpace(cmdLine)
		args := strings.Split(cmdLine, " ")
		if len(args) == 0 || args[0] == "" {
			continue
		}
		cmd := args[0]

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
			if len(args) < 2 {
				fmt.Println("Please provide port")
				continue
			}
			port, _ := strconv.Atoi(args[1])
			SetPort(cfg, port)
		case "setkey":
			if len(args) < 2 {
				fmt.Println("Please provide secret")
				continue
			}
			SetSecret(cfg, args[1])
		case "exit", "quit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println(Lang[CurrentLang]["usage"])
		}
	}
}
