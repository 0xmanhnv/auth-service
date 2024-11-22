package handler

import (
	"auth-service/internal/model"
	"auth-service/pkg/response"
	"auth-service/pkg/utils"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// LoginWithTelegramHandler godoc
// @Summary Login with Telegram
// @Description Handles login via Telegram authentication
// @Tags auth
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Param username query string true "Username"
// @Param first_name query string true "First Name"
// @Param auth_date query int true "Authentication Date"
// @Param photo_url query string false "Photo URL"
// @Success 200 {object} map[string]interface{} "Login successful"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /auth/telegram [get]
func (uHandler *AuthHandler) LoginWithTelegramHandler(c *gin.Context) {
	// Lấy thông tin từ query parameters
	userID := c.Query("id")
	username := c.Query("username")
	firstName := c.Query("first_name")
	authDate := c.Query("auth_date")
	photoUrl := c.Query("photo_url")

	// Kiểm tra thời gian xác thực (không quá 5 phút)
	if time.Now().Unix()-utils.ParseInt(authDate) > 300 {
		c.JSON(
			http.StatusUnauthorized,
			response.Response{
				Status:  "error",
				Message: "Login failed",
				Error:   "Expired auth date",
			},
		)
		return
	}

	// Get all query parameters
	params := c.Request.URL.Query()
	// So sánh mã băm
	if !checkTelegramAuthorization(params) {
		c.JSON(
			http.StatusUnauthorized,
			response.Response{
				Status:  "error",
				Message: "Login failed",
				Error:   "Invalid hash",
			},
		)
		return
	}

	// Nếu tất cả đều hợp lệ, xử lý đăng nhập
	user := model.User{
		Username:       username,
		FirstName:      firstName,
		PhotoUrl:       photoUrl,
		ChatTelegramID: int64(utils.ParseInt(userID)), // Chuyển đổi userID,
	}
	token, err := uHandler.AuthService.LoginWithTelegram(c, user)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			response.Response{
				Status:  "error",
				Message: "Login failed",
				Error:   err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.Response{
			Status:  "success",
			Message: "Login successful",
			Data: gin.H{
				"username": username,
				"token":    token,
			},
		},
	)
}

// id, first_name, last_name, username, photo_url, auth_date and hash
func checkTelegramAuthorization(params url.Values) bool {
	// Create a hash of the secret key
	keyHash := sha256.New()
	keyHash.Write([]byte(os.Getenv("TELEGRAM_BOT_TOKEN")))
	secretkey := keyHash.Sum(nil)

	// Check the hash
	var checkparams []string
	for k, v := range params {
		if k != "hash" {
			checkparams = append(checkparams, fmt.Sprintf("%s=%s", k, v[0]))
		}
	}

	// Sort the parameters
	sort.Strings(checkparams)

	// Join and create the hash
	checkString := strings.Join(checkparams, "\n")
	hash := hmac.New(sha256.New, secretkey)
	hash.Write([]byte(checkString))
	hashstr := hex.EncodeToString(hash.Sum(nil))

	return hashstr == params["hash"][0]
}
