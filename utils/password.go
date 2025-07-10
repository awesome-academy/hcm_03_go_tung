package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword mã hóa mật khẩu
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash kiểm tra mật khẩu có đúng không
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
