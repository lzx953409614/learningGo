package main

//原始net/http实现web端http接口
import (
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	result, err := io.WriteString(w, "hello this is a go http server")
	if err != nil {
		log.Printf("自定义go http server失败！err=%s",err)
	}
	log.Printf("自定义go http server 成功！result=%d",result)
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
