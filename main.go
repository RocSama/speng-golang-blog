package main

import (
	"log"
	"net/http"
	"speng-golang-blog/common"
	"speng-golang-blog/router"
)

func init() {
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "localhost:9090",
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
