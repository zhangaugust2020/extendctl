# extendctl

### 打包Linux运行
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
 
#打包
go build -o extendctl .
chmod +x extendctl
```
#### 主要命令
```
# 获取Pod信息
extendctl get pod
```
```
# 通过Pod名字查找部署该Pod的Deployment
extendctl podOwnDeploy [Pod名字]
```
```
# 重启Deployment部署的Pod
extendctl restart [Deployment名字]
```
```
# 查看Deployment部署的版本
extendctl version [Deployment名字]
```
```
# 获取Node信息
extendctl get node
```
```
# 获取Service信息
extendctl get svc
```