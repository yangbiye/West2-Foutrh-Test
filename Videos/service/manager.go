package service

import (
	"errors"
	"fmt"

	models "github.com/videos/Models"
	util "github.com/videos/pkg/util/jwt"
)

// 管理员登录
func RegisterManager(password string, email string) (uint, string, error) {
	if password == "" || email == "" {
		return 0, "", errors.New("请完整填写邮箱和密码")
	}
	manager, result := models.RegisterManager(password, email)
	if result.Error != nil {
		return 0, "", result.Error
	}
	//自动生成用户ID
	token, err := util.GenerateToken2(manager.ID)

	if err != nil {
		return 0, "", err
	}

	return manager.ID, token, nil
}

// 管理员登录
func LoginManager(managerID uint, pwd string) (string, error) {
	manager, res := models.LoginManager(managerID)

	if res.Error != nil {
		return "", res.Error
	}
	fmt.Println(pwd)
	fmt.Println(manager.Password)
	if pwd != manager.Password {
		return "", errors.New("密码不正确")
	}

	token, err := util.GenerateToken2(manager.ID)

	if err != nil {
		return "", err
	}
	return token, nil
}

// 审核视频
func PassVideo(videoID uint) error {
	res := models.ReleaseVideo(videoID)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

// 拉黑用户
func LockUser(userID uint) error {
	res := models.LockUser(userID)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

// 拉白用户
func DeBockUser(userID uint) error {
	res := models.DeBockUser(userID)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

// 封号用户
func DeleteUser(userID uint) error {
	//删除用户
	res := models.DeleteUser(userID)
	if res.Error != nil {
		return res.Error
	}
	//删除视频
	res1 := models.DeleteVideoByUserID(userID)
	if res1.Error != nil {
		return res1.Error
	}
	//删除弹幕
	res2 := models.DeleteBarrageByUserID(userID)
	if res2.Error != nil {
		return res2.Error
	}
	//删除评论
	res3 := models.DeleteCommentByUserID(userID)
	if res3.Error != nil {
		return res3.Error
	}
	//删除收藏
	res4 := models.DeleteFavoriteByUserID(userID)
	if res4.Error != nil {
		return res4.Error
	}
	return nil
}

func DeleteComment(commentID uint) error {
	comment, res := models.SearchComment(commentID)
	if res.Error != nil {
		return res.Error
	}
	if comment.RootParentID == 0 {
		res2 := models.DeleteRootParentCom(commentID)
		if res2.Error != nil {
			return res2.Error
		} else {
			return nil
		}
	} else { //若该评论不是根评论，就删除这个分支
		res3 := DeleteAllChildrenCom(commentID)
		if res3 != nil {
			return res3
		} else {
			return nil
		}

	}
}

// func DeleteAll(commentID uint) *gorm.DB {
// 	comments, res := models.SearchCommentByParID(commentID)
// 	if res.Error != nil {
// 		return res
// 	}
// 	if len(comments) == 0 {
// 		//delete 自己
// 		return nil
// 	} else {
// 		for i := 0; i < len(comments); i++ {
// 			DeleteAll(comments[i].ID)
// 		}
// 		return nil
// 	}
// }

func DeleteAllChildrenCom(commentID uint) error {
	//查询所有评论id
	childrenID, res := models.SearchAllChildrenID(commentID)
	if res.Error != nil {
		return res.Error
	}
	//删除评论及子孙
	for _, id := range childrenID {
		if err := DeleteAllChildrenCom(id); err != nil {
			return err
		}
		if err := models.DeleteComment(id).Error; err != nil {
			return err
		}
	}
	//删除自身
	err := models.DeleteComment(commentID).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
