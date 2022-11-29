##Linux 环境部署

###本地编译
```
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux
go build -o asm-platform-api
```
### 本地启动
```
asm-platform-api D:\goProjects\go\src\asm-platform-api\etc\config.ini

指定config文件路径地址
```