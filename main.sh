#!/bin/bash
###
 # @作者: idchubkk
 # @日期: 2022-07-01
 # @说明: 一键安装 / 管理 MTProto 代理
 # @Telegram: https://t.me/idchubkk
###

PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

# 颜色定义
red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

# 必须使用 root
[[ $EUID -ne 0 ]] && echo -e "[${red}错误${plain}] 请使用 ROOT 用户运行此脚本！" && exit 1

download_file(){
	echo "正在检测系统架构..."

	bit=`uname -m`
	if [[ ${bit} = "x86_64" ]]; then
		bit="amd64"
    elif [[ ${bit} = "aarch64" ]]; then
        bit="arm64"
    else
	    bit="386"
    fi

    last_version=$(curl -Ls "https://api.github.com/repos/9seconds/mtg/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    if [[ ! -n "$last_version" ]]; then
        echo -e "${red}获取最新版本失败，可能是 Github API 访问频率限制，请稍后再试。${plain}"
        exit 1
    fi
    echo -e "检测到 mtg 最新版本: ${last_version}，开始安装..."
    version=$(echo ${last_version} | sed 's/v//g')
    wget -N --no-check-certificate -O mtg-${version}-linux-${bit}.tar.gz https://github.com/9seconds/mtg/releases/download/${last_version}/mtg-${version}-linux-${bit}.tar.gz
    if [[ ! -f "mtg-${version}-linux-${bit}.tar.gz" ]]; then
        echo -e "${red}下载 mtg-${version}-linux-${bit}.tar.gz 失败，请重试。${plain}"
        exit 1
    fi
    tar -xzf mtg-${version}-linux-${bit}.tar.gz
    mv mtg-${version}-linux-${bit}/mtg /usr/bin/mtg
    rm -f mtg-${version}-linux-${bit}.tar.gz
    rm -rf mtg-${version}-linux-${bit}
    chmod +x /usr/bin/mtg
    echo -e "mtg 安装成功，开始配置..."
}

configure_mtg(){
    echo -e "正在配置 mtg..."
    wget -N --no-check-certificate -O /etc/mtg.toml https://raw.githubusercontent.com/missuo/MTProxy/main/mtg.toml
    
    echo ""
    read -p "请输入伪装域名（默认 itunes.apple.com）: " domain
	[ -z "${domain}" ] && domain="itunes.apple.com"

	echo ""
    read -p "请输入监听端口（默认 8443）: " port
	[ -z "${port}" ] && port="8443"

    secret=$(mtg generate-secret --hex $domain)
    
    echo "正在写入配置..."

    sed -i "s/secret.*/secret = \"${secret}\"/g" /etc/mtg.toml
    sed -i "s/bind-to.*/bind-to = \"0.0.0.0:${port}\"/g" /etc/mtg.toml

    echo "配置完成，开始设置 systemctl..."
}

configure_systemctl(){
    echo -e "正在配置 systemctl 服务..."
    wget -N --no-check-certificate -O /etc/systemd/system/mtg.service https://raw.githubusercontent.com/missuo/MTProxy/main/mtg.service
    systemctl enable mtg
    systemctl start mtg
    echo "关闭防火墙..."
    systemctl disable firewalld
    systemctl stop firewalld
    ufw disable
    echo "MTProxy 启动成功！"
    echo ""
    public_ip=$(curl -s ipv4.ip.sb)
    subscription_config="tg://proxy?server=${public_ip}&port=${port}&secret=${secret}"
    subscription_link="https://t.me/proxy?server=${public_ip}&port=${port}&secret=${secret}"
    echo -e "=== 你的 MTProto 代理信息 ==="
    echo -e "IP: ${public_ip}"
    echo -e "端口: ${port}"
    echo -e "密钥: ${secret}"
    echo -e "\nTelegram 点击链接：\n${subscription_config}\n${subscription_link}"
}

change_port(){
    read -p "请输入新的端口（默认 8443）: " port
	[ -z "${port}" ] && port="8443"
    sed -i "s/bind-to.*/bind-to = \"0.0.0.0:${port}\"/g" /etc/mtg.toml
    echo "正在重启 MTProxy..."
    systemctl restart mtg
    echo "端口修改成功并已重启！"
}

change_secret(){
    echo -e "注意：修改 Secret 可能导致客户端无法连接！"
    read -p "请输入新的 Secret（留空则使用默认域名生成）: " secret
	[ -z "${secret}" ] && secret="$(mtg generate-secret --hex itunes.apple.com)"
    sed -i "s/secret.*/secret = \"${secret}\"/g" /etc/mtg.toml
    echo "Secret 修改成功！"
    echo "正在重启 MTProxy..."
    systemctl restart mtg
    echo "重启完成！"
}

update_mtg(){
    echo -e "正在更新 mtg..."
    download_file
    echo "更新完成，正在重启 MTProxy..."
    systemctl restart mtg
    echo "重启完成！"
}

start_menu() {
    clear
    echo -e " MTProto 一键管理脚本 v2
---- 作者: Vincent | 中文化 by ChatGPT ----
 ${green} 1.${plain} 安装 MTProxy
 ${green} 2.${plain} 卸载 MTProxy
————————————
 ${green} 3.${plain} 启动 MTProxy
 ${green} 4.${plain} 停止 MTProxy
 ${green} 5.${plain} 重启 MTProxy
 ${green} 6.${plain} 修改监听端口
 ${green} 7.${plain} 修改 Secret
 ${green} 8.${plain} 更新 MTProxy
————————————
 ${green} 0.${plain} 退出
————————————" && echo

	read -e -p "请输入数字 [0-8]: " num
	case "$num" in
    1)
		download_file
        configure_mtg
        configure_systemctl
		;;
    2)
        echo "正在卸载 MTProxy..."
        systemctl stop mtg
        systemctl disable mtg
        rm -rf /usr/bin/mtg
        rm -rf /etc/mtg.toml
        rm -rf /etc/systemd/system/mtg.service
        echo "卸载完成！"
        ;;
    3) 
        echo "正在启动 MTProxy..."
        systemctl start mtg
        systemctl enable mtg
        echo "启动完成！"
        ;;
    4) 
        echo "正在停止 MTProxy..."
        systemctl stop mtg
        systemctl disable mtg
        echo "已停止！"
        ;;
    5)  
        echo "正在重启 MTProxy..."
        systemctl restart mtg
        echo "重启完成！"
        ;;
    6) 
        change_port
        ;;
    7)
        change_secret
        ;;
    8)
        update_mtg
        ;;
    0) exit 0
        ;;
    *) echo -e "${red}请输入正确的数字 [0-8]${plain}"
        ;;
    esac
}
start_menu
