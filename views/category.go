package views

import (
	"errors"
	"net/http"
	"speng-golang-blog/common"
	"speng-golang-blog/service"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	category := common.Template.Category
	//取参数
	//r都能取到什么参数？
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		//errors.New和log.print(err)的区别
		category.WriteError(w, errors.New("不能识别此路径"))
		return
	}
	//r取到了什么
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10

	categoryResponse, err := service.GetPostByCategoryId(cId, page, pageSize)
	if err != nil {
		category.WriteError(w, err)
		return
	}

	category.WriteData(w, categoryResponse)
}
