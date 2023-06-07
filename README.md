# mtg
MTProto协议介绍
MTProto协议是 Telegram 为了对抗网络封锁开发的专用代理（MTProxy）协议，目前全平台的 TG 客户端中都支持MTProto协议和MTProxy代理。有了MTProxy代理，即使没有VPN或者其他代理的情况下，也能顺畅访问TG。

本文介绍一键搭建Telegram的MTProto代理。

一键搭建Telegram的MTProto代理
第一步，请准备一台的非大陆VPS，请确保VPS的IP未被封锁，如遇封锁请及时提交工单更换IP，通常服务器商家都支持两小时内免费更换。VPS系统可以选择centos、ubuntu、Debian，我这里使用的是ubuntu 20.04

第二步，SSH登录到服务器

第三步，执行下面的命令一键搭建Telegram的MTProto代理：

# CentOS/AliyunOS/AMI系统

yum install -y curl
bash <(curl -sL https://idchubkk.github.io/idchub/mtg.sh)
# Debian/Ubuntu系统

apt install -y curl
bash <(curl -sL https://idchubkk.github.io/idchub/mtg.sh)
紧接着会出现如下菜单：


因为我们是首次安装所以选择 1，回车，按照提示输入一个端口号并回车（端口号随便设置，不和其他软件冲突即可，如果开启了防火墙则需要放行该端口）。

安装成功后，会输出如下信息：


第三步 打开telegram软件，参考 配置Telegram走SS/SSR/V2ray/trojan代理 的操作添加自定义代理，选择MTPROTO，将一键脚本输出的IP、端口和密钥填上去，点击保存：

接下来，就可以在不开启代理/VPN的情况下使用TG客户端了。

也可以直接将生成的连接复制到telegram打开

注意事项
目前MTProto已经发展到第三代，已经不建议使用V2ray内置的MTProto来搭建
本脚本使用了 9seconds 的docker镜像搭建；
因为docker访问外网需求，因此禁用了VPS的防火墙。如果你的VPS用于网站等重要业务，不建议使用本脚本搭建；
如果有国内VPS，建议使用 中转，防止被封；


