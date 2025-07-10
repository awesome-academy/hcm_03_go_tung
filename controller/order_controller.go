package controllers

import (
	"foods-drinks-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService}
}

func (c *OrderController) GetOrderHistory(ctx *gin.Context) {
	userIDStr := ctx.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	history, err := c.orderService.GetOrderHistory(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, history)
}

func (c *OrderController) PlaceOrder(ctx *gin.Context) {
	userIDStr := ctx.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	var req services.PlaceOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := c.orderService.PlaceOrder(userID, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Order placed successfully"})
}
