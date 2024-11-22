package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	if origin == "" {
		return
	}

	if c.Request.Method == http.MethodOptions {
		h := c.Writer.Header()
		h.Set("Access-Control-Allow-Origin", origin)
		h.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,HEAD")
		h.Set("Access-Control-Allow-Headers", "authorization,content-type")
		h.Set("Access-Control-Max-Age", "86400")
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	h := c.Writer.Header()
	h.Set("Access-Control-Allow-Origin", origin)
}
