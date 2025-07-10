package services

import (
	"fmt"
	"hcm_03_go_tung/config"
	"hcm_03_go_tung/models"
	"hcm_03_go_tung/utils"
	"net/smtp"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Regex kiểm tra định dạng email hợp lệ
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Gửi email reset mật khẩu
func SendResetEmail(email string) error {

	if !strings.Contains(email, "@") {
		email += "@gmail.com"
	}

	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}

	// Tìm người dùng theo email
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return fmt.Errorf("email not found")
	}

	// Tạo token reset
	token, _ := utils.GenerateResetToken(user.ID)
	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", token)

	from := "your_email@gmail.com"
	pass := "your_email_password"
	to := email
	msg := "Subject: Reset Password\n\nClick this link to reset your password:\n" + resetLink

	return smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
}

func ResetPassword(token, newPassword string) error {
	claims, err := utils.ParseResetToken(token)
	if err != nil {
		return fmt.Errorf("invalid or expired token")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	// Cập nhật mật khẩu mới trong DB
	return config.DB.Model(&models.User{}).
		Where("id = ?", claims.UserID).
		Update("password", string(hashed)).Error
}
