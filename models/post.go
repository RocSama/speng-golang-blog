package models

import (
	"html/template"
	"speng-golang-blog/config"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`         // 文章ID
	Title      string    `json:"title"`       // 文章ID
	Slug       string    `json:"slug"`        // 自定也页面 path
	Content    string    `json:"content"`     // 文章的html
	Markdown   string    `json:"markdown"`    // 文章的Markdown
	CategoryId int       `json:"category_id"` //分类id
	UserId     int       `json:"user_id"`     //用户id
	ViewCount  int       `json:"view_count"`  //查看次数
	Type       int       `json:"type"`        //文章类型 0 普通，1 自定义文章
	CreateAt   time.Time `json:"create_at"`   // 创建时间
	UpdateAt   time.Time `json:"update_at"`   // 更新时间
}

type PostMore struct {
	Pid          int           `json:"pid"`          // 文章ID
	Title        string        `json:"title"`        // 文章ID
	Slug         string        `json:"slug"`         // 自定也页面 path
	Content      template.HTML `json:"content"`      // 文章的html
	CategoryId   int           `json:"categoryId"`   // 文章的Markdown
	CategoryName string        `json:"categoryName"` // 分类名
	UserId       int           `json:"userId"`       // 用户id
	UserName     string        `json:"userName"`     // 用户名
	ViewCount    int           `json:"viewCount"`    // 查看次数
	Type         int           `json:"type"`         // 文章类型 0 普通，1 自定义文章
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}

type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"` // 文章ID
	Title string `orm:"title" json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
