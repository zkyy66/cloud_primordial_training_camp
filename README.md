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

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    app: http-server
  name: httpsvc
spec:
  ports:
    - port: 80
      protocol: TCP
      targetPort: 8080
  selector:
    app: http-server
  type: ClusterIP

```

安装ingress，这个过程并未采用领教的安装方式， 因为无法从k8s.gcr上拉取成功，故采用官网方式把文件下载到本地之后(deploy.yaml->ingress-nginx-deploy.yaml)，把关于k8s.gcr拉取镜像地址进行修改

```yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.2.0/deploy/static/provider/cloud/deploy.yaml
```

```
          #把相关关于gcr修改的部分贴出来，因为整个yaml文件太长
          #image: k8s.gcr.io/ingress-nginx/controller:v1.2.0@sha256:d8196e3bc1e72547c5dec66d6556c0ff92a23f6d0919b206be170bc90d5f9185
          image: registry.cn-hangzhou.aliyuncs.com/google_containers/nginx-ingress-controller:v1.2.0
          
                    #image: k8s.gcr.io/ingress-nginx/kube-webhook-certgen:v1.1.1@sha256:64d8c73dca984af206adf9d6d7e46aa550362b1d7a01f3a0a91b20cc67868660
          image: registry.cn-hangzhou.aliyuncs.com/google_containers/kube-webhook-certgen:v1.1.1
          
                    #image: k8s.gcr.io/ingress-nginx/kube-webhook-certgen:v1.1.1@sha256:64d8c73dca984af206adf9d6d7e46aa550362b1d7a01f3a0a91b20cc67868660
          image: registry.cn-hangzhou.aliyuncs.com/google_containers/kube-webhook-certgen:v1.1.1
```

安装Metallb

```yaml
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/namespace.yaml 

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/metallb.yaml
```

创建名为config且命名空间为 metallb-system的configMap

```
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: metallb-system
  name: config
data:
  config: |
    address-pools:
    - name: default
      protocol: layer2
      #采用2层负载，对于目前本人来说较为简单，address为ip地址范围
      addresses:
      - 192.168.50.27-192.168.50.250

```

关于证书在此不一一说明

在最后创建一个kind为Ingress的yaml

```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    cert-manager.io/issuer: letsencrypt-prod
  name: http-server
spec:
  ingressClassName: nginx
  rules:
    - host: http-server.51.cafe
      http:
        paths:
          - backend:
              service:
                name: httpsvc
                port:
                  number: 80
            path: /
            pathType: Prefix
  tls:
    - hosts:
        - http-server.51.cafe
      secretName: http-server

```

## 模块10作业：通过Grafana和Prometheus监控httpServer
修改了入口文件main.go中调用的client_and_server.ClientRequest()
```go
//路径：exercises/module2/client_and_server/client.go
func ClientRequest() {
	//http.HandleFunc("/", HandleClientRequest)
	//http.HandleFunc("/healthz", HandleHealth)

	metrics.Register()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndexRequest)
	mux.HandleFunc("/index", handleIndexRequest)
	mux.HandleFunc("/client", HandleClientRequest)
	mux.HandleFunc("/healthz", HandleHealth)

	mux.Handle("/metrics", promhttp.Handler())

	errInfo := http.ListenAndServe(":8080", mux)
	if errInfo != nil {
		log.Fatalf("Error %s\n", errInfo)
	}
}
//新增了方法
func handleIndexRequest(w http.ResponseWriter, r *http.Request) {

	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))

	user := r.URL.Query().Get("index")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello %s\n", user))
	} else {
		io.WriteString(w, "hello 请正确输入URL\n")
	}

	io.WriteString(w, "***********请求详情*************")
	log.Printf("响应的多少时间：%d ms", delay)
}

新增了目录和文件module10/metrics下的metrics.go负责向prometheus注册
```
通过孟老师教程中关于GrafanaDashboardJson文件构建图形化
```json
{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "target": {
          "limit": 100,
          "matchAny": false,
          "tags": [],
          "type": "dashboard"
        },
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "gnetId": null,
  "graphTooltip": 0,
  "id": 4,
  "links": [],
  "panels": [
    {
      "datasource": "Prometheus",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 9,
        "w": 12,
        "x": 0,
        "y": 0
      },
      "id": 2,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single"
        }
      },
      "targets": [
        {
          "exemplar": true,
          "expr": "histogram_quantile(0.95, sum(rate(httpserver_execution_latency_seconds_bucket[5m])) by (le))",
          "interval": "",
          "legendFormat": "",
          "refId": "A"
        },
        {
          "exemplar": true,
          "expr": "histogram_quantile(0.90, sum(rate(httpserver_execution_latency_seconds_bucket[5m])) by (le))",
          "hide": false,
          "interval": "",
          "legendFormat": "",
          "refId": "B"
        },
        {
          "exemplar": true,
          "expr": "histogram_quantile(0.50, sum(rate(httpserver_execution_latency_seconds_bucket[5m])) by (le))",
          "hide": false,
          "interval": "",
          "legendFormat": "",
          "refId": "C"
        }
      ],
      "title": "Response Latency by Percentile",
      "type": "timeseries"
    }
  ],
  "refresh": "",
  "schemaVersion": 30,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-1m",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "Http Server Latency",
  "uid": "mWgwgx5nz",
  "version": 2
}
```
#### 图片展示结果如下
![e7d5114a4363abe405bbbbf31e147ee](https://user-images.githubusercontent.com/756021/175799871-bba6d65c-2dc6-459e-8f78-5d35461acbe1.png)
![37f787ca9c915d60a89d76fe129d108](https://user-images.githubusercontent.com/756021/175799877-0f325c41-12f7-4b8b-9721-061b2235896d.png)
![f31c283da5fc94799fdbd691f8668f7](https://user-images.githubusercontent.com/756021/175799878-92adea71-073a-422f-80b9-763745576844.png)
