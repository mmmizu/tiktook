package service

import (
	"errors"
	"github.com/bwmarrin/snowflake"
	"tiktook/models"
	"tiktook/util"
)

func Register(name, password string) (user *models.User, err error) {
	result := models.GetDb().Where("name = ?", name).Limit(1).Find(&user)
	if result.RowsAffected != 0 {
		err = errors.New("user already exists")
		return
	}
	user.Name = name
	// 通过bcrypt加密密码
	user.Password = util.BcryptHash(password)
	// 通过雪花算法生成唯一ID
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		err = errors.New("register fail")
	}
	user.Id = int64(n.Generate())
	err = models.GetDb().Create(user).Error
	return
}
func Login(name, password string) (user *models.User, err error) {
	result := models.GetDb().Where("name = ?", name).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		err = errors.New("username does not exist")
		return
	}
	if ok := util.BcryptCheck(password, user.Password); !ok {
		err = errors.New("wrong password")
		return
	}
	return
}

func UserInfoByUserId(id int64) (user *models.User, err error) {
	result := models.GetDb().Where("id = ?", id).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		err = errors.New("user does not exist")
		return
	}
	return
}
