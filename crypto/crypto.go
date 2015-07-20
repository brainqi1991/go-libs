package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

// Md5 对字符串进行加密，并返回加密后的string(Golang官方返回的是[]byte不方便操作)
func Md5(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha1(data string) string {
	h := sha1.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Base64EncodeString 编码data
func Base64EncodeString(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64DecodeString 解码data
func Base64DecodeString(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(b), err
}
