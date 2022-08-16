package main

//原始net/http实现web端http接口
import (
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello this is a go http server")
}

func createHttpServer() {
	log.Println("创建一个go http server 开始。。。。")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatal("Error:", err)
	}
	log.Println("创建一个go http server 结束。。。。")
}
