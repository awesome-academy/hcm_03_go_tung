package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var resetSecret = []byte("super_secret_reset_key")

// ResetClaim định nghĩa payload JWT khi reset password
type ResetClaim struct {
	UserID uuid.UUID // ID của người dùng cần reset password
	jwt.StandardClaims
}

// Tạo token reset có hạn trong 30 phút
func GenerateResetToken(userID uuid.UUID) (string, error) {
	claims := ResetClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
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

	if claims, ok := token.Claims.(*ResetClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
