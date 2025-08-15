package main

var Lang = map[string]map[string]string{
	"zh": {
		"install": "安装 MTProxy",
		"start":   "启动代理",
		"stop":    "停止代理",
		"restart": "重启代理",
		"setport": "修改端口",
		"setkey":  "修改密钥",
		"usage":   "用法: mtmanager <install|start|stop|restart|setport|setkey> [参数]",
		"link":    "代理链接示例: tg://proxy?server=%s&port=%d&secret=%s",
	},
	"en": {
		"install": "Install MTProxy",
		"start":   "Start Proxy",
		"stop":    "Stop Proxy",
		"restart": "Restart Proxy",
		"setport": "Set Port",
		"setkey":  "Set Secret",
		"usage":   "Usage: mtmanager <install|start|stop|restart|setport|setkey> [args]",
		"link":    "Proxy link example: tg://proxy?server=%s&port=%d&secret=%s",
	},
}

var CurrentLang = "zh"
