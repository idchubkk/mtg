package main

import (
	"crypto/aes"
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

// AES-IGE 加密
func AESIGEEncrypt(block cipher.Block, plaintext, iv []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	xPrev := iv[:block.BlockSize()]
	yPrev := iv[block.BlockSize():]

	for i := 0; i < len(plaintext); i += block.BlockSize() {
		for j := 0; j < block.BlockSize(); j++ {
			plaintext[i+j] ^= yPrev[j]
		}
		block.Encrypt(ciphertext[i:i+block.BlockSize()], plaintext[i:i+block.BlockSize()])
		for j := 0; j < block.BlockSize(); j++ {
			ciphertext[i+j] ^= xPrev[j]
		}
		copy(xPrev, ciphertext[i:i+block.BlockSize()])
		copy(yPrev, plaintext[i:i+block.BlockSize()])
	}
	return ciphertext
}

// AES-IGE 解密
func AESIGEDecrypt(block cipher.Block, ciphertext, iv []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	xPrev := iv[:block.BlockSize()]
	yPrev := iv[block.BlockSize():]

	for i := 0; i < len(ciphertext); i += block.BlockSize() {
		for j := 0; j < block.BlockSize(); j++ {
			ciphertext[i+j] ^= xPrev[j]
		}
		block.Decrypt(plaintext[i:i+block.BlockSize()], ciphertext[i:i+block.BlockSize()])
		for j := 0; j < block.BlockSize(); j++ {
			plaintext[i+j] ^= yPrev[j]
		}
		copy(xPrev, ciphertext[i:i+block.BlockSize()])
		copy(yPrev, ciphertext[i:i+block.BlockSize()])
	}
	return plaintext
}
