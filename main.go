package main

import (
	"cloud_primordial_training_camp/exercises/module2/client_and_server"
)

func main() {
	client_and_server.ClientRequest()
}

//func main() {
//	//fmt.Println("start 1.1")
//	//handleStr := []string{"I", "am", "stupid", "and", "weak"}
//	//tmpRes := module1.HandleString(handleStr)
//	//fmt.Println(tmpRes)
//	//
//	//res := module1.HandleSliceString(handleStr)
//	//fmt.Println(res)
//	//
//	//result := module1.HandleSliceStringTwo(handleStr)
//	//fmt.Println("two:", result)
//	//fmt.Println("end 1.1")
//	//module8.GraceMain()
//}

//func httpDemoOne() {
//	http.HandleFunc("/", rootHandle)
//	http.ListenAndServe(":8080", nil)
//}
//func rootHandle(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("aaaa")
//	//w.WriteHeader(http.StatusOK)
//}
//func httpMux() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/aaa", rootHandle)
//	port := 8080
//	portstr := ":" + strconv.Itoa(port)
//	//ser := &http.Server{
//	//	Addr:    portstr,
//	//	Handler: mux,
//	//}
//	time.Sleep(time.Second * 3)
//	http.ListenAndServe(portstr, mux)
//}
