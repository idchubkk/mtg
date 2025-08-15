package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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

func generateSecret() string {
	b := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i++ {
		b[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", b)
}

func getPublicIP() string {
	cmd := exec.Command("curl", "-s", "https://api.ipify.org")
	out, err := cmd.Output()
	if err != nil {
		return "YOUR_SERVER_IP"
	}
	return string(out)
}

func InstallMTProxy() {
	fmt.Println(Lang[CurrentLang]["install"])

	if _, err := os.Stat("MTProxy"); os.IsNotExist(err) {
		if _, err := os.Stat("mtproxy-src.tar.gz"); os.IsNotExist(err) {
			fmt.Println("Downloading MTProxy source code...")
			cmd := exec.Command("wget", "-O", "mtproxy-src.tar.gz", "https://github.com/telegram-mtproxy/MTProxy/archive/refs/heads/master.tar.gz")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("下载失败，请检查网络:", err)
				return
			}
		}

		fmt.Println("Extracting MTProxy source...")
		cmd := exec.Command("tar", "-xzf", "mtproxy-src.tar.gz")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

		os.Rename("MTProxy-master", "MTProxy")
	}

	fmt.Println("Compiling MTProxy...")
	cmd := exec.Command("make", "-C", "MTProxy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	fmt.Println("MTProxy 安装完成，可执行文件路径: MTProxy/objs/bin/mtproto-proxy")
}

func UninstallMTProxy() {
	fmt.Println(Lang[CurrentLang]["uninstall"])
	StopProxy()
	os.RemoveAll("MTProxy")
	fmt.Println("卸载完成")
}

func StartProxy(cfg *Config) {
	if cfg.Secret == "" {
		cfg.Secret = generateSecret()
		SaveConfig(cfg)
	}
	cmd := exec.Command(cfg.Path,
		"-p", strconv.Itoa(cfg.Port),
		"-S", cfg.Secret,
		"--aes-pwd", "proxy-secret", "/dev/null",
		"-M", "1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	fmt.Println(Lang[CurrentLang]["start"])
	fmt.Printf(Lang[CurrentLang]["link"]+"\n", getPublicIP(), cfg.Port, cfg.Secret)
}

func StopProxy() {
	exec.Command("pkill", "-f", "mtproto-proxy").Run()
	fmt.Println(Lang[CurrentLang]["stop"])
}

func RestartProxy(cfg *Config) {
	StopProxy()
	StartProxy(cfg)
	fmt.Println(Lang[CurrentLang]["restart"])
}

func SetPort(cfg *Config, port int) {
	cfg.Port = port
	SaveConfig(cfg)
	fmt.Println(Lang[CurrentLang]["setport"], port)
	RestartProxy(cfg)
}

func SetSecret(cfg *Config, secret string) {
	cfg.Secret = secret
	SaveConfig(cfg)
	fmt.Println(Lang[CurrentLang]["setkey"], secret)
	RestartProxy(cfg)
}

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
