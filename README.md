# cloud_primordial_training_camp
# 云原生
## 基础之GoLang
### 模块三作业说明
```text
* 在模块二中main.go文件在exercises中
* 而模块三中把main.go移动和exercises同层级下，原因在Dockerfile编译中会提示go.mod缺少，当然也可以通过RUN CGO_ENABLED=0 GOOS=linux解决
* 作为一个项目来说main.go相当于入口文件
```
```shell
 1:执行镜像命令
    docker pull zkyy66/http_serverv
2 :启动命令
    docker run -it -d -p 自定义端口:（容器）8080 --name 自定义名称 镜像ID
```

