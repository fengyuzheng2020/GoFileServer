# 使用官方的Go镜像作为基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app

# 将本地的文件添加到工作目录中
COPY . .

# 构建Go应用程序
RUN go build -o main .

# 暴露端口
EXPOSE 8080

# 设置默认值，可通过环境变量覆盖
ENV PORT=8080
ENV UPLOAD_DIR=/app/uploads

# 运行应用程序，接受环境变量设置端口和上传目录
CMD ["./main", "-port", "$PORT", "-upload-dir", "$UPLOAD_DIR"]


# docker build -t go-file-server .
# docker run -p 8081:8080 -e UPLOAD_DIR=/path/to/custom/uploads go-file-server