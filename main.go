package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"math/rand"
	"encoding/json"
)

type Config struct {
	Port   int    `json:"port"`
	Secret string `json:"secret"`
	Path   string `json:"path"`
}

var Lang = map[string]map[string]string{
	"zh": {
		"langSelect": "选择语言:",
		"menu":       "\n操作菜单:\n1: 安装 MTProxy\n2: 卸载 MTProxy\n3: 启动代理\n4: 停止代理\n5: 重启代理\n6: 修改端口\n7: 修改密钥\n8: 退出",
		"install":    "安装 MTProxy",
		"uninstall":  "卸载 MTProxy",
		"start":      "启动代理",
		"stop":       "停止代理",
		"restart":    "重启代理",
		"setport":    "修改端口",
		"setkey":     "修改密钥",
		"exit":       "退出",
		"enterPort":  "请输入端口号:",
		"enterKey":   "请输入密钥:",
		"link":       "代理链接示例: tg://proxy?server=%s&port=%d&secret=%s",
	},
	"en": {
		"langSelect": "Select language:",
		"menu":       "\nMenu:\n1: Install MTProxy\n2: Uninstall MTProxy\n3: Start Proxy\n4: Stop Proxy\n5: Restart Proxy\n6: Set Port\n7: Set Secret\n8: Exit",
		"install":    "Install MTProxy",
		"uninstall":  "Uninstall MTProxy",
		"start":      "Start Proxy",
		"stop":       "Stop Proxy",
		"restart":    "Restart Proxy",
		"setport":    "Set Port",
		"setkey":     "Set Secret",
		"exit":       "Exit",
		"enterPort":  "Enter port number:",
		"enterKey":   "Enter secret:",
		"link":       "Proxy link example: tg://proxy?server=%s&port=%d&secret=%s",
	},
	"jp": {
		"langSelect": "言語を選択:",
		"menu":       "\nメニュー:\n1: MTProxy をインストール\n2: MTProxy をアンインストール\n3: プロキシを開始\n4: プロキシを停止\n5: プロキシを再起動\n6: ポートを変更\n7: 秘密鍵を変更\n8: 終了",
		"install":    "MTProxy をインストール",
		"uninstall":  "MTProxy をアンインストール",
		"start":      "プロキシを開始",
		"stop":       "プロキシを停止",
		"restart":    "プロキシを再起動",
		"setport":    "ポートを変更",
		"setkey":     "秘密鍵を変更",
		"exit":       "終了",
		"enterPort":  "ポート番号を入力:",
		"enterKey":   "秘密鍵を入力:",
		"link":       "プロキシリンク例: tg://proxy?server=%s&port=%d&secret=%s",
	},
	"pt": {
		"langSelect": "Selecione a língua:",
		"menu":       "\nMenu:\n1: Instalar MTProxy\n2: Desinstalar MTProxy\n3: Iniciar Proxy\n4: Parar Proxy\n5: Reiniciar Proxy\n6: Alterar Porta\n7: Alterar Chave\n8: Sair",
		"install":    "Instalar MTProxy",
		"uninstall":  "Desinstalar MTProxy",
		"start":      "Iniciar Proxy",
		"stop":       "Parar Proxy",
		"restart":    "Reiniciar Proxy",
		"setport":    "Alterar Porta",
		"setkey":     "Alterar Chave",
		"exit":       "Sair",
		"enterPort":  "Digite o número da porta:",
		"enterKey":   "Digite a chave:",
		"link":       "Exemplo de link do proxy: tg://proxy?server=%s&port=%d&secret=%s",
	},
}

var CurrentLang = "en"
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

// 以下为各操作函数
func InstallMTProxy() {
	fmt.Println(Lang[CurrentLang]["install"])
	if _, err := os.Stat("MTProxy"); os.IsNotExist(err) {
		if _, err := os.Stat("mtproxy-src.tar.gz"); err == nil {
			cmd := exec.Command("tar", "-xzf", "mtproxy-src.tar.gz")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else if _, err := os.Stat("MTProxy.zip"); err == nil {
			cmd := exec.Command("unzip", "MTProxy.zip")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			fmt.Println("未找到源码包 mtproxy-src.tar.gz 或 MTProxy.zip，请放入同目录")
			return
		}
	}
	cmd := exec.Command("make", "-C", "MTProxy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	fmt.Println("安装完成")
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

	// 主循环菜单
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
			fmt.Print(Lang[CurrentLang]["enter
