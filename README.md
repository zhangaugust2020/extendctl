# extendctl

### 打包Linux运行
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
 
#打包
go build -o extendctl .
```
