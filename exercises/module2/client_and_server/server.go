/**
 * @Date 2022/4/22
 * @Name server
 * @VariableName
**/
/**
写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
*/
package client_and_server

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func HandleClientRequest(resHeader http.ResponseWriter, r *http.Request) {
	user, paths, versionName := handleEnv()
	log.Printf("用户%s，路径：%s，用户版本：%s\n", user, paths, versionName)
	resHeader.Header().Set("USER", user)
	resHeader.Header().Set("GO_PATHS", paths)
	resHeader.Header().Set("VERSION_NAME", versionName)

	requestHeader := r.Header
	for index, val := range requestHeader {
		for _, value := range val {
			resHeader.Header().Set(index, value)
		}
		//if index == "Accept" {
		//	resHeader.Header().Set(index, val[0])
		//}

	}
	resHeader.WriteHeader(http.StatusOK)
	log.Printf("打印Accetp的value %s\n", resHeader.Header().Get("Accept"))
	log.Printf("当前您请求地址：%s\n", r.RequestURI)
	log.Printf("当前您请求地址：%s\n", r.RemoteAddr)
	log.Printf("当前请求HOST：%s\n", r.Host)
}
func handleEnv() (x, y, z string) {
	x = os.Getenv("USER")
	y = os.Getenv("GOPATH")
	os.Setenv("VERSION_NAME", "YAOYUAN")
	z = os.Getenv("VERSION_NAME")
	return
}
func HandleHealth(resHeader http.ResponseWriter, r *http.Request) {
	io.WriteString(resHeader, strconv.Itoa(http.StatusOK))
}
