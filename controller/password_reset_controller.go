package controllers

import (
    "net/http"
    "hcm_03_go_tung/services"
    "github.com/gin-gonic/gin"
)


func ForgotPassword(c *gin.Context) {
    var req struct {
        Email string `json:"email"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
        return
    }

    err := services.SendResetEmail(req.Email)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Check your email to reset password"})
}


func ResetPassword(c *gin.Context) {
    var req struct {
        Token       string `json:"token"`
        NewPassword string `json:"new_password"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    err := services.ResetPassword(req.Token, req.NewPassword)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}