package controllers

import (
	"foods-drinks-app/services"
	"foods-drinks-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartController struct {
	cartService services.CartService
}

func NewCartController(cartService services.CartService) *CartController {
	return &CartController{cartService: cartService}
}

// GetCartHandler hiển thị giỏ hàng
func (c *CartController) GetCartHandler(ctx *gin.Context) {
	userIDStr := ctx.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidUserID)})
		return
	}

	cartItems, err := c.cartService.GetCart(userID)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.GetMessage(lang, utils.ErrorFailedToUpdateCart)})
		return
	}

	ctx.JSON(http.StatusOK, cartItems)
}

// AddToCartHandler thêm sản phẩm vào giỏ hàng
func (c *CartController) AddToCartHandler(ctx *gin.Context) {
	userIDStr := ctx.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidUserID)})
		return
	}

	var req struct {
		ProductID uuid.UUID `json:"product_id"`
		Quantity  int64     `json:"quantity"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidRequestBody)})
		return
	}

	if err := c.cartService.AddToCart(userID, req.ProductID, req.Quantity); err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorFailedToAddToCart)})
		return
	}

	lang := ctx.MustGet("language").(utils.Language)
	ctx.JSON(http.StatusOK, gin.H{"message": utils.GetMessage(lang, utils.SuccessProductAddedToCart)})
}

// UpdateCartItemHandler cập nhật số lượng sản phẩm trong giỏ hàng
func (c *CartController) UpdateCartItemHandler(ctx *gin.Context) {
	userIDStr := ctx.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidUserID)})
		return
	}

	var req struct {
		ProductID uuid.UUID `json:"product_id"`
		Quantity  int64     `json:"quantity"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidRequestBody)})
		return
	}

	if err := c.cartService.UpdateCartItem(userID, req.ProductID, req.Quantity); err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorFailedToUpdateCart)})
		return
	}

	lang := ctx.MustGet("language").(utils.Language)
	ctx.JSON(http.StatusOK, gin.H{"message": utils.GetMessage(lang, utils.SuccessCartItemUpdated)})
}

// RemoveFromCartHandler xóa sản phẩm khỏi giỏ hàng
func (c *CartController) RemoveFromCartHandler(ctx *gin.Context) {
	userIDStr := ctx.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidUserID)})
		return
	}

	productIDStr := ctx.Param("product_id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidProductID)})
		return
	}

	if err := c.cartService.RemoveFromCart(userID, productID); err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorFailedToRemoveFromCart)})
		return
	}

	lang := ctx.MustGet("language").(utils.Language)
	ctx.JSON(http.StatusOK, gin.H{"message": utils.GetMessage(lang, utils.SuccessProductRemovedFromCart)})
}

// ClearCartHandler xóa toàn bộ giỏ hàng
func (c *CartController) ClearCartHandler(ctx *gin.Context) {
	userIDStr := ctx.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidUserID)})
		return
	}

	if err := c.cartService.ClearCart(userID); err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.GetMessage(lang, utils.ErrorFailedToClearCart)})
		return
	}

	lang := ctx.MustGet("language").(utils.Language)
	ctx.JSON(http.StatusOK, gin.H{"message": utils.GetMessage(lang, utils.SuccessCartCleared)})
}

// GetDeletedCartItemsHandler xem danh sách sản phẩm đã bị xóa
func (c *CartController) GetDeletedCartItemsHandler(ctx *gin.Context) {
	userIDStr := ctx.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidUserID)})
		return
	}

	deletedItems, err := c.cartService.GetDeletedCartItems(userID)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.GetMessage(lang, utils.ErrorFailedToUpdateCart)})
		return
	}

	ctx.JSON(http.StatusOK, deletedItems)
}

// RestoreCartItemHandler khôi phục sản phẩm đã bị xóa
func (c *CartController) RestoreCartItemHandler(ctx *gin.Context) {
	userIDStr := ctx.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidUserID)})
		return
	}

	productIDStr := ctx.Param("product_id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidProductID)})
		return
	}

	if err := c.cartService.RestoreCartItem(userID, productID); err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorFailedToRestoreCart)})
		return
	}

	lang := ctx.MustGet("language").(utils.Language)
	ctx.JSON(http.StatusOK, gin.H{"message": utils.GetMessage(lang, utils.SuccessProductRestored)})
}
