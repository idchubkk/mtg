package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"log"
	"net"
)

// 处理客户端连接
func HandleConnection(conn net.Conn, secret []byte) {
	defer conn.Close()
	fmt.Println("新连接:", conn.RemoteAddr())

	buf := make([]byte, 64)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("读取握手错误:", err)
		return
	}

	// 校验 secret
	if n < len(secret) || string(buf[:len(secret)]) != string(secret) {
		log.Println("secret 校验失败")
		return
	}

	// 生成 AES key/iv
	key := make([]byte, 32)
	iv := make([]byte, 32)
	_, _ = rand.Read(key)
	_, _ = rand.Read(iv)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("AES 创建失败:", err)
		return
	}

	// 返回握手响应
	conn.Write(append(secret, key...))
	conn.Write(iv)

	// 后续可通过 AESIGEEncrypt / AESIGEDecrypt 加密/解密数据
}

// 启动代理监听
func StartProxy(port int, secret []byte) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("监听端口失败:", err)
	}
	fmt.Println("MTProxy 监听端口:", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept 错误:", err)
			continue
		}
		go HandleConnection(conn, secret)
	}
}
