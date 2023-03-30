package API

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/videos/middleware"
	"github.com/videos/service"
)

// 用户注册
func RegisterUser(c *gin.Context) {
	email := c.Query("email")
	pwd := c.Query("password")

	uid, token, err := service.RegisterUser(pwd, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":     "ok",
		"user_id": uid,
		"token":   token,
	})
}

// 用户登录
func LoginUser(c *gin.Context) {
	user_ID := c.Query("user_id")
	password := c.Query("password")
	userID, err := strconv.ParseUint(user_ID, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	user, err := service.SearchUSer(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if user.InBlackList {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户已被拉黑",
		})
		return
	}
	token, err := service.LoginUser(uint(userID), password)

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

// 修改个人资料
func SetInfo(c *gin.Context) {
	action := c.Query("action")
	userName := c.Query("userName")
	introduction := c.Query("introduction")
	place := c.Query("place")
	passWord := c.Query("passWord")
	switch action {
	case "0": //修改用户名
		if userName == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "userName为空",
			})
			return
		}

		err := service.SetUserName(middleware.UserID, userName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg2": err.Error(),
			})
			return
		}

	case "1": //修改简介
		if introduction == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "introduction为空",
			})
			return
		}

		err := service.SetIntroduction(middleware.UserID, introduction)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}

	case "2": //修改所在地
		if place == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "place为空",
			})
			return
		}

		err := service.SetPlace(middleware.UserID, place)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}

	case "3": //修改密码
		if passWord == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "passWord为空",
			})
			return
		}

		err := service.SetPassword(middleware.UserID, passWord)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}

	case "4": //修改头像
		//上传头像
		file, err := c.FormFile("fi")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg1": err.Error(),
			})
			return
		} else {
			dst := fmt.Sprintf("./%s", file.Filename)
			err1 := c.SaveUploadedFile(file, dst)
			if err1 != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "文件未保存",
				})
				return
			}
			err := service.SetPicture(middleware.UserID, dst)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": err.Error(),
				})
				return
			}
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "action错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}

// 展示个人资料
func Show(c *gin.Context) {
	userID := middleware.UserID
	user, err := service.ShowUserInfo(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg:": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":           user.ID,
		"email":        user.Email,
		"userName":     user.UserName,
		"introduction": user.Introduction,
		"place":        user.Place,
	})
}

// 搜索视频
func SearchVideo(c *gin.Context) {
	action := c.Query("action")
	_year := c.Query("year")
	Type := c.Query("Type")
	_month := c.Query("month")
	_day := c.Query("day")
	switch action {
	case "0": //按年份搜索
		if _year == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "year为空",
			})
			return
		}
		//保存搜索记录
		err := service.CreateRecord(_year)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		year, err := strconv.ParseInt(_year, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		videos, err := service.SearchVideoYear(int(year))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		for i := 0; i < len(videos); i++ {
			c.JSON(http.StatusOK, gin.H{
				"Video:": videos[i].ID,
			})
		}
	case "1": //按类别搜索
		if Type == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "year为空",
			})
			return
		}
		err := service.CreateRecord(Type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		videos, err := service.SearchVideoType(Type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		for i := 0; i < len(videos); i++ {
			c.JSON(http.StatusOK, gin.H{
				"Video:": videos[i].ID,
			})
		}
	case "2": //按发布时间搜索
		if _month == "" || _day == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "数据不完整",
			})
			return
		}
		err := service.CreateRecord(_month + "" + _day)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		month, err1 := strconv.ParseInt(_month, 10, 32)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err1.Error(),
			})
			return
		}
		day, err2 := strconv.ParseInt(_day, 10, 32)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err2.Error(),
			})
			return
		}
		videos, err3 := service.SearchVideoMandD(time.Month(month), int(day))
		if err3 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err3.Error(),
			})
			return
		}
		for i := 0; i < len(videos); i++ {
			c.JSON(http.StatusOK, gin.H{
				"Video:": videos[i].ID,
			})
		}
	}
}

// 搜索用户
func SearchUser(c *gin.Context) {
	place := c.Query("place")
	if place == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "place为空",
		})
		return
	}
	err := service.CreateRecord(place)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	users, err := service.SearchUserPlace(place)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err:": err.Error(),
		})
		return
	}
	var i int
	for i = 0; i < len(users); i++ {
		c.JSON(http.StatusOK, gin.H{
			"user": users[i].ID,
		})
	}
}

// 展示历史搜索记录
func ShowRecord(c *gin.Context) {
	records, err := service.SearchRecord()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err:": err.Error(),
		})
		return
	}
	for i := 0; i < len(records); i++ {
		c.JSON(http.StatusOK, gin.H{
			"record": records[i].Rec,
		})
	}
}

// 多条件筛选视频
func SelectVideo(c *gin.Context) {
	action := c.Query("action")
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "title为空",
		})
		return
	}
	err := service.CreateRecord(title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	switch action {
	case "0": //点赞量
		videos, err := service.SearchVideoThumbupOrd(title)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		var i int
		for i = 0; i < len(videos); i++ {
			c.JSON(http.StatusOK, gin.H{
				"user": videos[i].ID,
			})
		}
	case "1": //时间先后
		videos, err := service.SearchVideoTimeOrd(title)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		var i int
		for i = 0; i < len(videos); i++ {
			c.JSON(http.StatusOK, gin.H{
				"user": videos[i].ID,
			})
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "action错误",
		})
		return
	}
}
