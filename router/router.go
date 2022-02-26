package router

import (
	"net/http"
	"speng-golang-blog/api"
	"speng-golang-blog/views"
)

func Router() {
	//1.页面 2.数据 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
