package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

// 校验用户输入的明文密码与数据库中查询得到的密码是否匹配
func CheckPassword(password string, password_digest string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_digest), []byte(password))
	if err == nil {
		return true
	} else {
		return false
	}
}

// 生成bcrypt加密后的密码
func HasSecurePassword(password string) string {
	password_digest, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(password_digest)
}
