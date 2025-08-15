package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// 生成 16 字节随机密钥
func generateSecret() string {
	b := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i++ {
		b[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", b)
}

// 获取公网 IP
func getPublicIP() string {
	cmd := exec.Command("curl", "-s", "https://api.ipify.org")
	out, err := cmd.Output()
	if err != nil {
		return "YOUR_SERVER_IP"
	}
	return string(out)
}

// 安装 MTProxy
func InstallMTProxy() {
	fmt.Println(Lang[CurrentLang]["install"])

	if _, err := os.Stat("MTProxy"); os.IsNotExist(err) {
		if _, err := os.Stat("mtproxy-master.zip"); os.IsNotExist(err) {
			fmt.Println("Downloading MTProxy source code...")
			cmd := exec.Command("wget", "-O", "mtproxy-master.zip", "https://github.com/TelegramMessenger/MTProxy/archive/refs/heads/master.zip")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("下载失败，请检查网络:", err)
				return
			}
		}

		fmt.Println("Extracting MTProxy source...")
		cmd := exec.Command("unzip", "-q", "mtproxy-master.zip")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("解压失败，请安装 unzip 工具:", err)
			return
		}

		os.Rename("MTProxy-master", "MTProxy")
	}

	fmt.Println("Compiling MTProxy...")
	cmd := exec.Command("make", "-C", "MTProxy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	fmt.Println("MTProxy 安装完成，可执行文件路径: MTProxy/objs/bin/mtproto-proxy")
}

// 卸载 MTProxy
func UninstallMTProxy() {
	fmt.Println(Lang[CurrentLang]["uninstall"])
	StopProxy()
	os.RemoveAll("MTProxy")
	fmt.Println("卸载完成")
}

// 启动 MTProxy
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

// 停止 MTProxy
func StopProxy() {
	exec.Command("pkill", "-f", "mtproto-proxy").Run()
	fmt.Println(Lang[CurrentLang]["stop"])
}

// 重启 MTProxy
func RestartProxy(cfg *Config) {
	StopProxy()
	StartProxy(cfg)
	fmt.Println(Lang[CurrentLang]["restart"])
}

// 修改端口并重启
func SetPort(cfg *Config, port int) {
	cfg.Port = port
	SaveConfig(cfg)
	fmt.Println(Lang[CurrentLang]["setport"], port)
	RestartProxy(cfg)
}

// 修改密钥并重启
func SetSecret(cfg *Config, secret string) {
	cfg.Secret = secret
	SaveConfig(cfg)
	fmt.Println(Lang[CurrentLang]["setkey"], secret)
	RestartProxy(cfg)
}
