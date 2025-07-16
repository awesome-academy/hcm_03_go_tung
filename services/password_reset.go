package services

import (
	"fmt"
	"hcm_03_go_tung/config"
	"hcm_03_go_tung/models"
	"hcm_03_go_tung/utils"
	"net/smtp"
	"os"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/message"
	"github.com/nicksnyder/go-i18n/v2/i18n" // Import the i18n package
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
	token, err := utils.GenerateResetToken(user.ID)
	if err != nil {
		return fmt.Errorf(i18n.T("failed_to_generate_reset_token"), err)
	}

	// Lấy hostname từ config hoặc biến môi trường
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = config.AppConfig.Hostname // Assuming you have a Hostname field in your config
	}
	resetLink := fmt.Sprintf("http://%s/reset-password?token=%s", hostname, token)

	from := os.Getenv("EMAIL_ADDRESS")
	pass := os.Getenv("EMAIL_PASSWORD")
	to := email
	msg := "Subject: " + i18n.T("reset_password_subject") + "\n\n" + i18n.T("reset_password_message", resetLink)

	return smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
}
