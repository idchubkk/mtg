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
	fmt.Println("Select language / 选择语言 / 言語を選択 / Selecione a língua:")
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
		CurrentLang = "en"
	}

	cfg := LoadConfig()

	for {
		fmt.Println(Lang[CurrentLang]["menu"])
		fmt.Print("Choice: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			InstallMTProxy()
		case 2:
			UninstallMTProxy()
		case 3:
			StartProxy(cfg)
		case 4:
			StopProxy()
		case 5:
			RestartProxy(cfg)
		case 6:
			fmt.Print(Lang[CurrentLang]["enterPort"])
			portStr, _ := reader.ReadString('\n')
			portStr = strings.TrimSpace(portStr)
			port, _ := strconv.Atoi(portStr)
			SetPort(cfg, port)
		case 7:
			fmt.Print(Lang[CurrentLang]["enterKey"])
			secret, _ := reader.ReadString('\n')
			secret = strings.TrimSpace(secret)
			SetSecret(cfg, secret)
		case 8:
			fmt.Println(Lang[CurrentLang]["exit"])
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
