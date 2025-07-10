package routes

import (
	controllers "foods-drinks-app/controller"
	"foods-drinks-app/middleware"
	"foods-drinks-app/repository"
	"foods-drinks-app/services"

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
	authGroup.Use(middleware.LanguageMiddleware())
	authGroup.Use(middleware.JWTAuthMiddleware())

	{
		authGroup.GET("/", profileController.GetProfileHandler)
		authGroup.PUT("/", profileController.UpdateProfileHandler)
		authGroup.PUT("/change-password", profileController.ChangePasswordHandler)
	}
}

func SetupOrderRoutes(r *gin.Engine, db *gorm.DB) {
	orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	authGroup := r.Group("/api/orders")
	authGroup.Use(middleware.LanguageMiddleware())
	authGroup.Use(middleware.JWTAuthMiddleware())
	{
		authGroup.GET("/history", orderController.GetOrderHistory)
		authGroup.POST("/", orderController.PlaceOrder)
	}
}

func SetupCartRoutes(r *gin.Engine, db *gorm.DB) {
	cartRepo := repository.NewCartRepository(db)
	cartService := services.NewCartService(cartRepo)
	cartController := controllers.NewCartController(cartService)

	authGroup := r.Group("/api/cart")
	authGroup.Use(middleware.LanguageMiddleware())
	authGroup.Use(middleware.JWTAuthMiddleware())
	{
		authGroup.GET("/", cartController.GetCartHandler)
		authGroup.POST("/add", cartController.AddToCartHandler)
		authGroup.PUT("/update", cartController.UpdateCartItemHandler)
		authGroup.DELETE("/remove/:product_id", cartController.RemoveFromCartHandler)
		authGroup.DELETE("/clear", cartController.ClearCartHandler)
		authGroup.GET("/deleted", cartController.GetDeletedCartItemsHandler)
		authGroup.POST("/restore/:product_id", cartController.RestoreCartItemHandler)
	}
}
