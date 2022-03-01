package dao

import (
	"log"
	"speng-golang-blog/models"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreatAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 查询出错", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("SELECT name from blog_category WHERE cid=?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	row.Scan(&categoryName)
	return categoryName
}
