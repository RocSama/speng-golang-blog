package dao

import "log"

func GetUserNameById(userId int) string {
	row := DB.QueryRow("SELECT user_name from blog_user WHERE uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	row.Scan(&userName)
	return userName
}
