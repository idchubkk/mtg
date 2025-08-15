package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func InstallMTProxy() {
	fmt.Println(Lang[CurrentLang]["install"])
	if _, err := os.Stat("MTProxy"); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", "https://github.com/telegram-mtproxy/MTProxy.git")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	cmd := exec.Command("make", "-C", "MTProxy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func generateSecret() string {
	b := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i++ {
		b[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", b)
}

func StartProxy(cfg *Config) {
	if cfg.Secret == "" {
		cfg.Secret = generateSecret()
	}
	SaveConfig("config.json", cfg)
	cmd := exec.Command(cfg.Path,
		"-p", fmt.Sprintf("%d", cfg.Port),
		"-S", cfg.Secret,
		"--aes-pwd", "proxy-secret", "/dev/null",
		"-M", "1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Println("启动失败:", err)
		return
	}
	fmt.Println(Lang[CurrentLang]["start"])
	fmt.Printf(Lang[CurrentLang]["link"]+"\n", getPublicIP(), cfg.Port, cfg.Secret)
}

func StopProxy() {
	cmd := exec.Command("pkill", "-f", "mtproto-proxy")
	cmd.Run()
	fmt.Println(Lang[CurrentLang]["stop"])
}

func RestartProxy(cfg *Config) {
	StopProxy()
	StartProxy(cfg)
	fmt.Println(Lang[CurrentLang]["restart"])
}

func SetPort(cfg *Config, port int) {
	cfg.Port = port
	SaveConfig("config.json", cfg)
	fmt.Println(Lang[CurrentLang]["setport"], port)
	RestartProxy(cfg)
}

func SetSecret(cfg *Config, secret string) {
	cfg.Secret = secret
	SaveConfig("config.json", cfg)
	fmt.Println(Lang[CurrentLang]["setkey"], secret)
	RestartProxy(cfg)
}

// 获取公网 IP（简单实现，可用 curl)
func getPublicIP() string {
	cmd := exec.Command("curl", "-s", "https://api.ipify.org")
	out, err := cmd.Output()
	if err != nil {
		return "YOUR_SERVER_IP"
	}
	return string(out)
}
