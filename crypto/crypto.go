package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// Md5 对字符串进行加密，并返回加密后的string(Golang官方返回的是[]byte不方便操作)
func Md5(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Sha1 return the string crpyto with Sha1 value
func Sha1(data string) string {
	h := sha1.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Sha256 return the string crypto with 256 value
func Sha256(data string) string {
	h := sha256.New()
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

// CheckPassword 校验用户输入的明文密码与数据库中查询得到的密码是否匹配
func CheckPassword(password string, passwordDigest string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDigest), []byte(password))
	if err == nil {
		return true
	}
	return false
}

// HasSecurePassword 生成bcrypt加密后的密码
func HasSecurePassword(password string) string {
	passwordDigest, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordDigest)
}
