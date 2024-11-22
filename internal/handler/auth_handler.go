package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// Import middleware
	"auth-service/internal/model"
	"auth-service/pkg/response"
)

type AuthHandler struct {
	AuthService AuthService
}

func NewAuthHandler(authSvc AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authSvc,
	}
}

// RegisterHandler godoc
// @Summary Register a new user
// @Description Handles user registration
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.User true "User information"
// @Success 201 {object} map[string]interface{} "User created"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 409 {object} map[string]interface{} "Conflict"
// @Router /auth/register [post]
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Response{
				Status:  "error",
				Message: "Invalid input",
				Error:   err.Error(),
				Data:    nil,
			},
		)
		return
	}
	if err := h.AuthService.Register(c, user); err != nil {
		c.JSON(
			http.StatusConflict,
			response.Response{
				Status:  "error",
				Message: "Invalid input",
				Error:   err.Error(),
				Data:    nil,
			},
		)
		return
	}
	c.JSON(
		http.StatusCreated,
		response.Response{
			Status:  "success",
			Message: "User created",
			Error:   "",
			Data:    nil,
		},
	)
}

// LoginHandler godoc
// @Summary Login a user
// @Description Handles user login
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.LoginRequest true "User credentials"
// @Success 200 {object} map[string]interface{} "Login successful"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /auth/login [post]
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var user model.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	token, err := h.AuthService.Login(c, user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			response.Response{
				Status:  "error",
				Message: "Unauthorized",
				Error:   err.Error(),
				Data:    nil,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		response.Response{
			Status:  "success",
			Message: "Login successful",
			Error:   "",
			Data: gin.H{
				"token": token,
			},
		},
	)
}

// ProtectedHandler godoc
// @Summary Access protected route
// @Description Handles requests to a protected route
// @Tags auth
// @Produce json
// @Success 200 {object} map[string]interface{} "Protected route access"
// @Router /auth/protected [get]
// @Security BearerAuth
func (h *AuthHandler) ProtectedHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		response.Response{
			Status:  "success",
			Message: "Protected route access",
			Error:   "",
			Data:    nil,
		},
	)
}
