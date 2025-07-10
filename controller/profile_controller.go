package controllers

import (
	"foods-drinks-app/services"
	"foods-drinks-app/utils"
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
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidUserID)})
		return
	}

	var request struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		lang := ctx.MustGet("language").(utils.Language)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.GetMessage(lang, utils.ErrorInvalidRequestBody)})
		return
	}

	lang := ctx.MustGet("language").(utils.Language)
	err = c.profileService.ChangePassword(userID, request.OldPassword, request.NewPassword, lang)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": utils.GetMessage(lang, utils.SuccessPasswordChanged)})
}
