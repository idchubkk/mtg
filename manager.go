package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

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
