package middleware

import (
	"net/http"
	"strings"

	"auth-service/internal/config"
	"auth-service/pkg/jwt"
	"auth-service/pkg/response"

	"github.com/gin-gonic/gin"
)

// ValidateToken là middleware để xác thực token JWT
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized,
				response.Response{
					Status:  "error",
					Message: "Unauthorized",
					Error:   "Authorization header is required",
					Data:    nil,
				},
			)
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.ValidateToken(tokenString, config.Cfg.JWTSecretKey)

		if err != nil || !token.Valid {
			c.JSON(
				http.StatusUnauthorized,
				response.Response{
					Status:  "error",
					Message: "Unauthorized",
					Error:   err.Error(),
					Data:    nil,
				},
			)
			c.Abort()
			return
		}

		c.Next()
	}
}
