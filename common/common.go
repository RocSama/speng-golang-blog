package common

import (
	"speng-golang-blog/config"
	"speng-golang-blog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template")
}
