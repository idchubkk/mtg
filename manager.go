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
