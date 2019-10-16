# 基于 golang-1.13 与 gin-gonic/gin 的 api 服务
> 使用 go modules 管理依赖

# Feature
|name|router|
|----|----|
|ping|http://127.0.0.1:65535/v1/ping|
|snowflake uuid generator|http://127.0.0.1:65535/v1/id|


# Usage
```shell script
go mod tidy
go run main.go
```

# Build
```shell script
go mod tidy
go build -o gofun main.go
```

# Notice
若依赖安装时网络错误，可添加代理：
```shell script
# export GOPROXY=https://mirrors.aliyun.com/goproxy/
export GOPROXY="https://goproxy.io"
export GOSUMDB=gosum.io+ce6e7565+AY5qEHUk/qmHc5btzW45JVoENfazw8LielDsaI+lEbq6
``` 

# Docs
使用 swagger 实现 API 文档自动生成
```shell script
go get -u github.com/swaggo/swag/cmd/swag
cd gofun
swag init
```
SwaggerUI: http://localhost:65535/swagger/index.html
