package service

import (
	"errors"

	models "github.com/videos/Models"
	util "github.com/videos/pkg/util/jwt"
)

// 用户注册
func RegisterUser(password string, email string) (uint, string, error) {
	if password == "" || email == "" {
		return 0, "", errors.New("请完整填写邮箱和密码")
	}
	user, result := models.RegisterUser(password, email)
	if result.Error != nil {
		return 0, "", result.Error
	}
	//自动生成用户ID
	token, err := util.GenerateToken1(user.ID)

	if err != nil {
		return 0, "", err
	}

	return user.ID, token, nil
}

// 用户登录
func LoginUser(userID uint, pwd string) (string, error) {
	user, res := models.LoginUser(userID)

	if res.Error != nil {
		return "", res.Error
	}

	if pwd != user.Password {
		return "", errors.New("密码不正确")
	}

	token, err := util.GenerateToken1(user.ID)

	if err != nil {
		return "", err
	}
	return token, nil
}

// 展示个人资料
func ShowUserInfo(userID uint) (models.User, error) {
	user, res := models.LoginUser(userID)
	var user1 models.User
	if res.Error != nil {
		return user1, res.Error
	} else {
		return user, nil
	}

}

// 设置用户名
func SetUserName(userID uint, userName string) error {
	res := models.SetUserName(userID, userName)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

// 设置简介
func SetIntroduction(userID uint, introduction string) error {
	res := models.SetIntroduction(userID, introduction)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

// 设置所在地
func SetPlace(userID uint, place string) error {
	res := models.SetPlace(userID, place)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

// 修改密码
func SetPassword(userID uint, password string) error {
	res := models.SetPassword(userID, password)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

// 修改头像
func SetPicture(userID uint, picture string) error {
	res := models.SetPicture(userID, picture)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

func SearchUSer(userID uint) (models.User, error) {
	var user models.User
	var user1 models.User
	user, err := models.SearchUSer(userID)
	if err.Error != nil {
		return user1, err.Error
	} else {
		return user, nil
	}
}

// 创建历史搜索记录
func CreateRecord(rec string) error {
	err := models.CreateRecord(rec)
	if err.Error != nil {
		return err.Error
	} else {
		return nil
	}
}
