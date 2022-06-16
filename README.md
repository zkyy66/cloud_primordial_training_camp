# cloud_primordial_training_camp
# 云原生

## GitHub地址

```tex
云原生地址：https://github.com/zkyy66/cloud_primordial_training_camp
```

## 基础之GoLang

### 模块三作业说明

```text
* 在模块二中main.go文件在exercises中，而exercises目录下的模块三则只有一个READER说明
* 而模块三中把main.go移动和exercises同层级下，原因在Dockerfile编译中会提示go.mod缺少，当然也可以通过RUN CGO_ENABLED=0 GOOS=linux解决
* 作为一个项目来说main.go相当于入口文件
```

#### 模块三步骤

```shell
0. 按照要求Dockfile要求是多阶段构建
1. 执行镜像pull命令：
    1. docker pull zkyy66/http_serverv
2. 启动命令：
    1. docker run -it -d -p 自定义端口:（容器）8080 --name 自定义名称 镜像ID
3. hub.docker地址：
    1. https://hub.docker.com/r/zkyy66/http_serverv
```

### 模块八作业第一部分
yaml文件在exercises/module8/http-server-deployment.yaml

参考资料：https://blog.csdn.net/weixin_39927378/article/details/111010625
```yaml
#configMap
apiVersion: v1
metadata:
  name: loglevel
kind: ConfigMap
data:
  httpport: "8080"
  loglevel: "info"

---

#Deployment方式部署
apiVersion: apps/v1
kind: Deployment
metadata:
  name: http-server-deploy
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: http-server
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: http-server
    spec:
      containers:
        - name: http-server-zkyy66
          image: docker.io/zkyy66/http_serverv:latest
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 8080
          lifecycle:
            postStart:
              exec:
                command: [ "/bin/sh", "-c", "echo Hello from the http server handler > /usr/message" ]
            preStop:
              exec:
                command: [ "/bin/sh","-c","ps -ef | grep monitor.go | grep grep -v | awk '{print $2}' | xargs kill" ]
          livenessProbe:
            httpGet:
              path: /healthz
              port: 8080
              scheme: HTTP
            initialDelaySeconds: 5
            periodSeconds: 10
            successThreshold: 2
            timeoutSeconds: 1
          resources:
            limits:
              cpu: 200m
              memory: 400Mi
            requests:
              cpu: 100m
              memory: 200Mi
          readinessProbe:
            httpGet:
              path: /healthz
              port: 8080
            initialDelaySeconds: 30
            periodSeconds: 5
            successThreshold: 2
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
      volumes:
        - name: loglevel
          configMap:
            name: loglevel
```
### 模块八作业第二部分说明

相关文件地址：exercises/module8/part2目录下

建立HTTPServer的svc

安装ingress，这个过程并未采用领教的安装方式， 因为无法从k8s.gcr上拉取成功，故采用官网方式把文件下载到本地之后(deploy.yaml->ingress-nginx-deploy.yaml)，把关于k8s.gcr拉取镜像地址进行修改

```yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.2.0/deploy/static/provider/cloud/deploy.yaml
```

安装Metallb

```yaml
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/namespace.yaml 

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/metallb.yaml
```

创建名为config且命名空间为 metallb-system的configMap

关于证书在此不一一说明，采用的是孟老师那种手动生成方式

在最后创建一个kind为Ingress的yaml
