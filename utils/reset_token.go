package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var resetSecret = []byte(os.Getenv("RESET_SECRET")) // Lấy giá trị từ biến môi trường

// ResetClaim định nghĩa payload JWT khi reset password
type ResetClaim struct {
	UserID uint // ID của người dùng cần reset password
	jwt.RegisteredClaims
}

// Tạo token reset có hạn trong 30 phút
func GenerateResetToken(userID uint) (string, error) {
	claims := ResetClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(resetSecret)
}

// Parse và xác thực token reset
func ParseResetToken(tokenStr string) (*ResetClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &ResetClaim{}, func(t *jwt.Token) (interface{}, error) {
		return resetSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*ResetClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
