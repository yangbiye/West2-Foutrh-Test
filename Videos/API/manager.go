package API

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/videos/service"
)

// 用户注册
func RegisterManager(c *gin.Context) {
	email := c.Query("email")
	pwd := c.Query("password")
	mid, token, err := service.RegisterManager(pwd, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":        "ok",
		"manager_id": mid,
		"token":      token,
	})
}

// 登录
func LoginManager(c *gin.Context) {
	manager_ID := c.Query("manager_id")
	password := c.Query("password")
	managerID, err := strconv.ParseUint(manager_ID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err:": err.Error(),
		})
		return
	}
	token, err := service.LoginManager(uint(managerID), password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"token": token,
	})
}

// 审核视频
func PassVideo(c *gin.Context) {
	video_ID := c.Query("videoID")
	videoID, err := strconv.ParseUint(video_ID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err1 := service.PassVideo(uint(videoID))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err1.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"videoID": videoID,
		"msg":     "视频审核成功",
	})
}

// 管理用户
func ManageUser(c *gin.Context) {
	action := c.Query("action")
	user_ID := c.Query("userID")
	if user_ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "userid为空",
		})
	}
	userID, err := strconv.ParseUint(user_ID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	switch action {
	case "0": //拉黑用户
		err := service.LockUser(uint(userID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userID": userID,
			"msg":    "拉黑成功",
		})
	case "1": //解封用户
		err := service.DeBockUser(uint(userID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userID": userID,
			"msg":    "解封成功",
		})
	case "2": //封号用户
		err := service.DeleteUser(uint(userID))
		fmt.Println(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userID": userID,
			"msg":    "已封号",
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "action错误",
		})
		return
	}

}

// 删除评论
func DeleteComment(c *gin.Context) {
	comment_ID := c.Query("commentID")
	if comment_ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "commentid为空",
		})
	}
	commentID, err := strconv.ParseUint(comment_ID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err1 := service.DeleteComment(uint(commentID))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err1.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "删除成功",
	})
}
