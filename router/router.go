package router

import (
	"net/http"
	"speng-golang-blog/api"
	"speng-golang-blog/views"
)

func Router() {
	//1.页面 2.数据 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	//category api
	http.HandleFunc("/c/", views.HTML.Category)
	//用户登录
	http.HandleFunc("/login", views.HTML.Login)
	//登录api
	http.HandleFunc("/api/v1/login", api.API.Login)
	//文章详情
	http.HandleFunc("/p/", views.HTML.Detail)
	//写文章
	http.HandleFunc("/writing", views.HTML.Writing)
	//保存文章
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	//归档
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	//搜索
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
}
