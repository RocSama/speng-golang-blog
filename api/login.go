package api

import (
	"net/http"
	"speng-golang-blog/common"
	"speng-golang-blog/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	//fmt.Printf(userName + " " + passwd)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
