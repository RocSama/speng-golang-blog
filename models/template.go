package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		w.Write([]byte("error"))
	}
}

func InitTemplate(templateDir string) HtmlTemplate {
	tp := readTemplate(
		[]string{"index", "category", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate
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

func readTemplate(templates []string, templateDir string) []TemplateBlog {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//拿到当前路径
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GerNextName, "date": Date})
		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println("解析模板出错：", err)
		}
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs
}
