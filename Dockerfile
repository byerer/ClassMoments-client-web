# 使用 golang 1.21 官方镜像作为构建阶段的基础镜像
FROM golang:1.21 as builder

# 设置工作目录为 /app
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct

# 将当前目录下的所有文件复制到工作目录
COPY . .

RUN go mod download
RUN go install github.com/google/wire/cmd/wire@latest
RUN wire ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/ClassMoments


# 第二阶段：运行阶段
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# 从构建器阶段复制编译后的可执行文件
COPY --from=builder /app/myapp .
CMD ["./myapp"]
