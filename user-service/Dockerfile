# user-service/Dockfile

# 使用 golang 官方镜像提供 Go 运行环境，并且命名为 buidler 以便后续引用
FROM golang:1.13-alpine as builder

# 启用 Go Modules 并设置 GOPROXY
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# 安装 git
RUN apk --no-cache add git
# 设置工作目录
WORKDIR /app/laracom-user-service

COPY . .

# 下载依赖
RUN go mod download
# 构建二进制文件，添加来一些额外参数以便可以在 Alpine 中运行它
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-user-service






FROM alpine:latest

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装相关软件
RUN apk update && apk add --no-cache bash ca-certificates

RUN mkdir /app
WORKDIR /app

# 这一步不再从宿主机拷贝二进制文件，而是从上一个阶段构建的 builder 容器中拉取
COPY --from=builder /app/laracom-user-service/laracom-user-service .


# CMD ["/usr/bin/supervisord","-c","/etc/supervisord.conf"]
CMD ["./laracom-user-service"]

