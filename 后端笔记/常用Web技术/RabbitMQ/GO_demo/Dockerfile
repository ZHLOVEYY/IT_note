# 使用官方 Go 镜像作为构建阶段
FROM golang:1.23 AS builder

WORKDIR /app

# 复制 go.mod 和 go.sum，下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建 Go 应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用轻量级镜像运行
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]