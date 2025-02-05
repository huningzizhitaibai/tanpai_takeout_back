package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func PasswordCrypto(password string) string {
	//使用SHA-256加密算法将密码进行加密然后存储到数据库提升安全性
	hash := sha256.Sum256([]byte(password))
	hash_password := hex.EncodeToString(hash[:])
	return hash_password
}
