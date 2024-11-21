package handler

import (
	"auth-service/internal/model"
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

func (uHandler *AuthHandler) LoginWithTelegramHandler(c *gin.Context) {
	// Lấy thông tin từ query parameters
	userID := c.Query("id")
	username := c.Query("username")
	firstName := c.Query("first_name")
	authDate := c.Query("auth_date")
	photoUrl := c.Query("photo_url")

	// Kiểm tra thời gian xác thực (không quá 5 phút)
	if time.Now().Unix()-utils.ParseInt(authDate) > 300 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Expired auth date"})
		return
	}

	// Get all query parameters
	params := c.Request.URL.Query()
	// So sánh mã băm
	if !checkTelegramAuthorization(params) {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "Invalid hash"},
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message":  "Login successful",
			"username": username,
			"token":    token,
		},
	)
}

// id, first_name, last_name, username, photo_url, auth_date and hash
func checkTelegramAuthorization(params url.Values) bool {
	keyHash := sha256.New()
	keyHash.Write([]byte(os.Getenv("TELEGRAM_BOT_TOKEN")))
	secretkey := keyHash.Sum(nil)

	var checkparams []string
	for k, v := range params {
		if k != "hash" {
			checkparams = append(checkparams, fmt.Sprintf("%s=%s", k, v[0]))
		}
	}
	sort.Strings(checkparams)
	checkString := strings.Join(checkparams, "\n")
	hash := hmac.New(sha256.New, secretkey)
	hash.Write([]byte(checkString))
	hashstr := hex.EncodeToString(hash.Sum(nil))

	return hashstr == params["hash"][0]
}
