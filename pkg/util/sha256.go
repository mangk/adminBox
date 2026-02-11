package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(input string) string {
	// 创建一个 SHA-256 哈希对象
	hasher := sha256.New()
	// 写入输入数据
	hasher.Write([]byte(input))
	// 计算哈希值
	hash := hasher.Sum(nil)
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hash)
}
