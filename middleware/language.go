package middleware

import (
	"foods-drinks-app/utils"

	"github.com/gin-gonic/gin"
)

// LanguageMiddleware xử lý ngôn ngữ từ Accept-Language header
func LanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		acceptLanguage := c.GetHeader("Accept-Language")
		lang := utils.GetLanguageFromHeader(acceptLanguage)
		c.Set("language", lang)
		c.Next()
	}
}
