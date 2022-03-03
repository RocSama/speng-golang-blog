package views

import (
	"net/http"
	"speng-golang-blog/common"
	"speng-golang-blog/config"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}
