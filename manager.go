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

	// 检查源码目录是否存在
	if _, err := os.Stat("MTProxy"); os.IsNotExist(err) {
		// 使用 wget 下载官方源码 tar.gz
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

		// 解压源码
		fmt.Println("Extracting MTProxy source...")
		cmd := exec.Command("tar", "-xzf", "mtproxy-src.tar.gz")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

		// 重命名解压后的目录为 MTProxy
		os.Rename("MTProxy-master", "MTProxy")
	}

	// 编译
	fmt.Println("Compiling MTProxy...")
	cmd := exec.Command("make", "-C", "MTProxy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("编译失败:", err)
		return
	}

	fmt.Println("MTProxy 安装完成，可执行文件路径: MTProxy/objs/bin/mtproto-proxy")
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
