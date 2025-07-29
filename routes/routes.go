package routes

import (
	"HCM_03_GO_TUNG/controllers"
	"HCM_03_GO_TUNG/middleware"
	"HCM_03_GO_TUNG/repository"
	"HCM_03_GO_TUNG/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProfileRoutes(r *gin.Engine, db *gorm.DB) {
	// Khởi tạo các dependency
	userRepo := repository.NewUserRepository(db)
	profileService := services.NewProfileService(userRepo)
	profileController := controllers.NewProfileController(profileService)

	// Nhóm các route cần xác thực
	authGroup := r.Group("/api/profile")
	authGroup.Use(middleware.JWTAuthMiddleware())

	{
		authGroup.GET("/", profileController.GetProfileHandler)
		authGroup.PUT("/", profileController.UpdateProfileHandler)
		authGroup.PUT("/change-password", profileController.ChangePasswordHandler)
	}
}