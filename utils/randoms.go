package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString tạo chuỗi ngẫu nhiên với độ dài cho trước
func RandString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RmEmail tạo email ngẫu nhiên theo định dạng <random>@example.com
func RandEmail() string {
	return RandString(10) + "@example.com"
}

// RandomPassword tạo mật khẩu ngẫu nhiên với ký tự đặc biệt
func RandPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*"
	var sb strings.Builder
	for i := 0; i < 12; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
