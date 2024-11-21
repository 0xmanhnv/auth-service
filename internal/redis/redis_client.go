package redis

import (
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client

func InitRedis() {

	redisDB, err := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 64) // 10 là cơ số, 64 là bit
	if err != nil {
		redisDB = 0
	}

	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"), // Mật khẩu nếu có
		DB:       int(redisDB),                // Sử dụng DB mặc định
	})

	// Kiểm tra kết nối
	if err := client.Ping(ctx).Err(); err != nil {
		panic(err)
	}
}

// Hàm để lấy client Redis
func GetClient() *redis.Client {
	return client
}
