package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// Import middleware
	"auth-service/internal/model"
	"auth-service/internal/service"
)

type AuthHandler struct {
	AuthService AuthService
}

func NewAuthHandler(authSvc *service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authSvc,
	}
}

// RegisterHandler xử lý yêu cầu đăng ký
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.AuthService.Register(c, user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// LoginHandler xử lý yêu cầu đăng nhập
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	token, err := h.AuthService.Login(c, user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ProtectedHandler xử lý yêu cầu đến route bảo vệ
func (h *AuthHandler) ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is a protected route"})
}
