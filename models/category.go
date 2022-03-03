package models

type Category struct {
	Cid      int
	Name     string
	CreatAt  string
	UpdateAt string
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
