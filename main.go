package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/helloworld", HelloWorld)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("消息接收失败: %v\n", err)
		return
	}
	log.Printf("消息已接收: %s\n", string(req))

	rsp := "Hello World!"
	io.WriteString(w, rsp)
	log.Printf("消息已返回: %s\n", string(rsp))
}
