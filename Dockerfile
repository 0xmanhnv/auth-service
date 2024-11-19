# Sử dụng image Go 1.23
FROM golang:1.23 AS builder

# Đặt thư mục làm việc
WORKDIR /app

# Sao chép go.mod và go.sum
COPY go.mod go.sum ./

# Tải các dependency
RUN go mod download

# Sao chép mã nguồn
COPY . .

# Biên dịch ứng dụng
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Bước 2: Tạo image cho ứng dụng
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
# Sao chép binary từ builder
COPY --from=builder /app/main .
COPY .env .

RUN chmod +x /app/main

EXPOSE 8080
# Chạy ứng dụng
CMD ["/app/main"]