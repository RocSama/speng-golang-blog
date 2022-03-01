package views

import (
	"errors"
	"log"
	"net/http"
	"speng-golang-blog/common"
	"speng-golang-blog/service"
	"strconv"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！！"))
	}
	index.WriteData(w, hr)

}

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	category := common.Template.Category
	//数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		category.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错", err)
		category.WriteError(w, errors.New("系统错误，请联系管理员！！"))
	}
	category.WriteData(w, hr)

}
