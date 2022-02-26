package main

import (
	"log"
	"net/http"
	"speng-golang-blog/config"
	"speng-golang-blog/models"
	"text/template"
	"time"
)

type IndexData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func IsODD(num int) bool {
	return num%2 == 0
}

func GerNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func index(w http.ResponseWriter, r *http.Request) {
	// var indexData IndexData
	// indexData.Name = "姜彪"
	// indexData.Description = "sb"
	t := template.New("index.html")
	//拿到当前路径
	path := config.Cfg.System.CurrentDir
	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GerNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println("解析模板出错：", err)
	}
	//页面上涉及到的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}

	t.Execute(w, hr)
}

func main() {
	server := http.Server{
		Addr: "localhost:9090",
	}
	http.HandleFunc("/", index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
