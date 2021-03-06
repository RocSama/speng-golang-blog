package dao

import (
	"fmt"
	"log"
	"speng-golang-blog/models"
)

// Pid        int       `json:"pid"`        // 文章ID
// 	Title      string    `json:"title"`      // 文章ID
// 	Slug       string    `json:"slug"`       // 自定也页面 path
// 	Content    string    `json:"content"`    // 文章的html
// 	Markdown   string    `json:"markdown"`   // 文章的Markdown
// 	CategoryId int       `json:"categoryId"` //分类id
// 	UserId     int       `json:"userId"`     //用户id
// 	ViewCount  int       `json:"viewCount"`  //查看次数
// 	Type       int       `json:"type"`       //文章类型 0 普通，1 自定义文章
// 	CreateAt   time.Time `json:"createAt"`   // 创建时间
// 	UpdateAt   time.Time `json:"updateAt"`   // 更新时间

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	rows.Scan(&count)
	return
}

func CountGetAllPostByCategoryId(cId int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id=?", cId)
	_ = rows.Scan(&count)
	return
}

func GetPostPageByCategoryId(cId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cId, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostById(pId int) (models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid=?", pId)
	var post models.Post
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return post, row.Err()
	}

	return post, nil
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println(err)
	}
}
func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into blog_post "+
		"(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	)
	if err != nil {
		log.Println(err)
	}
	pid, _ := ret.LastInsertId()
	post.Pid = int(pid)
}

func GetAllPost() ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post")
	if err != nil {
		log.Println("GetAllPost 查询出错", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			log.Println("GetAllPost 查询出错", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostSearch(condition string) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post where title like ?", "%"+condition+"%")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
