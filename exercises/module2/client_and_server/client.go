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
	"log"
	"net/http"
)

func ClientRequest() {
	//记录开始
	http.HandleFunc("/", HandleClientRequest)
	http.HandleFunc("/healthz", HandleHealth)
	errInfo := http.ListenAndServe(":8080", nil)
	if errInfo != nil {
		log.Fatalf("Error %s\n", errInfo)
	}
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
