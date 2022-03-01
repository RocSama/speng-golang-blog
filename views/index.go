package views

import (
	"errors"
	"log"
	"net/http"
	"speng-golang-blog/common"
	"speng-golang-blog/service"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	//数据库查询

	index := common.Template.Index
	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("Index获取数据出错", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！！"))
	}
	index.WriteData(w, hr)

}
