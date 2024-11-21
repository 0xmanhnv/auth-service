# Architech

```txt
my-go-microservices/
├── db/
│   ├── db.go
│   └── config.go
├── order-service/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   └── model/
│   ├── go.mod
│   └── go.sum
└── user-service/
    ├── cmd/
    │   └── main.go
    ├── internal/
    │   ├── handler/
    │   ├── service/
    │   ├── repository/
    │   └── model/
    ├── go.mod
    └── go.sum
```


OR

```txt
my-go-microservices/
├── order-service/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── handler/
│   │   │   └── handler.go
│   │   ├── service/
│   │   │   └── order_service.go
│   │   ├── repository/
│   │   │   └── order_repository.go
│   │   ├── model/
│   │   │   └── order.go
│   │   ├── config/
│   │   │   └── config.go
│   │   └── redis/
│   │       └── redis_client.go
│   ├── go.mod
│   └── go.sum
├── user-service/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── handler/
│   │   │   └── handler.go
│   │   ├── service/
│   │   │   └── user_service.go
│   │   ├── repository/
│   │   │   └── user_repository.go
│   │   ├── model/
│   │   │   └── user.go
│   │   ├── config/
│   │   │   └── config.go
│   │   └── socket/
│   │       └── socket_client.go
│   ├── go.mod
│   └── go.sum
└── api-gateway/
    ├── cmd/
    │   └── main.go
    ├── internal/
    │   └── handler/
    │       └── gateway.go
    ├── go.mod
    └── go.sum
```

Giải thích các thành phần
order-service: Dịch vụ quản lý đơn hàng.

repository/order_repository.go: Kết nối đến cơ sở dữ liệu (ví dụ: MySQL, PostgreSQL).
redis/redis_client.go: Kết nối đến Redis để lưu trữ dữ liệu tạm thời hoặc cache.
config/config.go: Quản lý cấu hình cho các kết nối đến cơ sở dữ liệu, Redis, và các dịch vụ khác.
user-service: Dịch vụ quản lý người dùng.

socket/socket_client.go: Kết nối đến các dịch vụ socket (ví dụ: WebSocket) để xử lý các yêu cầu thời gian thực.
api-gateway: Dịch vụ cổng API để điều phối các yêu cầu đến các dịch vụ khác.


https://github.com/bxcodec/go-clean-arch


## Gen docs

```bash
swag init -g .\cmd\main.go -o ./docs
```