package main

var Lang = map[string]map[string]string{
	"zh": {
		"install": "安装 MTProxy",
		"start":   "启动代理",
		"stop":    "停止代理",
		"restart": "重启代理",
		"setport": "修改端口",
		"setkey":  "修改密钥",
		"usage":   "可用命令: install | start | stop | restart | setport <端口> | setkey <密钥> | exit",
		"link":    "代理链接示例: tg://proxy?server=%s&port=%d&secret=%s",
	},
	"en": {
		"install": "Install MTProxy",
		"start":   "Start Proxy",
		"stop":    "Stop Proxy",
		"restart": "Restart Proxy",
		"setport": "Set Port",
		"setkey":  "Set Secret",
		"usage":   "Available commands: install | start | stop | restart | setport <port> | setkey <secret> | exit",
		"link":    "Proxy link example: tg://proxy?server=%s&port=%d&secret=%s",
	},
	"jp": {
		"install": "MTProxy をインストール",
		"start":   "プロキシを開始",
		"stop":    "プロキシを停止",
		"restart": "プロキシを再起動",
		"setport": "ポートを変更",
		"setkey":  "秘密鍵を変更",
		"usage":   "使用可能なコマンド: install | start | stop | restart | setport <port> | setkey <secret> | exit",
		"link":    "プロキシリンク例: tg://proxy?server=%s&port=%d&secret=%s",
	},
	"pt": {
		"install": "Instalar MTProxy",
		"start":   "Iniciar Proxy",
		"stop":    "Parar Proxy",
		"restart": "Reiniciar Proxy",
		"setport": "Alterar Porta",
		"setkey":  "Alterar Chave",
		"usage":   "Comandos disponíveis: install | start | stop | restart | setport <porta> | setkey <chave> | exit",
		"link":    "Exemplo de link do proxy: tg://proxy?server=%s&port=%d&secret=%s",
	},
}

var CurrentLang = "en"
