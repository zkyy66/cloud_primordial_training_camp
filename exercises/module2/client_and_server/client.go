/**
 * @Date 2022/4/22
 * @Name cliet
 * @VariableName
**/
/**
/**
写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
*/
package client_and_server

import (
	"cloud_primordial_training_camp/exercises/module10/metrics"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

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

func randInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//func HandleEnv() {
//	glog.V(2).Info(os.Environ())
//	envUser := os.Getenv("USER")
//	envPath := os.Getenv("GOPATH")
//	fmt.Printf("用户：%s ;gopath路：%s\n", envUser, envPath)
//}

//func Healthz() {
//	http.HandleFunc("/healthz", HandleHealth)
//	err := http.ListenAndServe("127.0.0.1:8085", nil)
//	if err != nil {
//		return
//	}
//}
