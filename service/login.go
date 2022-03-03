package service

import (
	"errors"
	"fmt"
	"speng-golang-blog/dao"
	"speng-golang-blog/models"
	"speng-golang-blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "mszlu")
	fmt.Println(passwd)
	user := dao.GerUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//生成token jwt生成令牌A.B.C
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token错误")
	}

	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar

	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
