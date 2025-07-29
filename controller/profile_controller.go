package controllers

import (
	"HCM_03_GO_TUNG/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProfileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) *ProfileController {
	return &ProfileController{profileService}
}

// GetProfileHandler xử lý request lấy thông tin hồ sơ
func (c *ProfileController) GetProfileHandler(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.GetString("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	profile, err := c.profileService.GetProfile(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

// UpdateProfileHandler xử lý request cập nhật hồ sơ
func (c *ProfileController) UpdateProfileHandler(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.GetString("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var updates map[string]interface{}
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	updatedProfile, err := c.profileService.UpdateProfile(userID, updates)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProfile)
}

// ChangePasswordHandler xử lý request đổi mật khẩu
func (c *ProfileController) ChangePasswordHandler(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.GetString("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var request struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = c.profileService.ChangePassword(userID, request.OldPassword, request.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})
}