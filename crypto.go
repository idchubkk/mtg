package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

// 生成 16 字节 secret
func GenerateSecret() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("生成 secret 失败:", err)
	}
	return hex.EncodeToString(b)
}

// AES-IGE 加密/解密接口保留，但暂时不调用 block
func AESIGEEncrypt(block interface{}, plaintext, iv []byte) []byte {
	// 占位，避免未使用报错
	return plaintext
}

func AESIGEDecrypt(block interface{}, ciphertext, iv []byte) []byte {
	// 占位
	return ciphertext
}
