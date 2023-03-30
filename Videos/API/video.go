package API

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/videos/middleware"
	"github.com/videos/service"
)

// 用户上传视频
func Upload(c *gin.Context) {
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
		userID := middleware.UserID
		Type := c.Query("Type")
		title := c.Query("title")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}

		err2 := service.UploadVideo(userID, dst, Type, title)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg2": err2.Error(),
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"msg": "上传成功",
	})
}

// 视频操作
func VideoOperate(c *gin.Context) {
	action := c.Query("action")
	userID := middleware.UserID
	video_ID := c.Query("videoID")
	comment_ID := c.Query("commentID")
	context := c.Query("context")
	switch action {
	case "0": //点赞
		if video_ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "videoID为空",
			})
			return
		}
		videoID, err := strconv.ParseUint(video_ID, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		video, err := service.SearchVideo(uint(videoID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		if video.InBlackList {
			c.JSON(http.StatusOK, gin.H{
				"msg": "视频未审核",
			})
			return
		}

		err1 := service.ThumbsUpVideo(uint(videoID))
		if err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err1.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "点赞成功",
		})
	case "1": //评论视频
		if video_ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "videoID为空",
			})
			return
		}
		videoID, err := strconv.ParseUint(video_ID, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		video, err := service.SearchVideo(uint(videoID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		if video.InBlackList {
			c.JSON(http.StatusOK, gin.H{
				"msg": "视频未审核",
			})
			return
		}
		if context == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "context为空",
			})
			return
		}
		err1 := service.CommentVideo(userID, uint(videoID), context)
		if err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err1.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "评论成功",
		})
	case "2": //评论评论
		if comment_ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "commentID为空",
			})
			return
		}
		commentID, err := strconv.ParseUint(comment_ID, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}

		if context == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "context为空",
			})
			return
		}
		err1 := service.CommentComment(userID, uint(commentID), context)
		if err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err1.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "评论成功",
		})
	case "3": //收藏
		if video_ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "videoID为空",
			})
			return
		}
		videoID, err := strconv.ParseUint(video_ID, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		video, err := service.SearchVideo(uint(videoID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		if video.InBlackList {
			c.JSON(http.StatusOK, gin.H{
				"msg": "视频未审核",
			})
			return
		}
		err1 := service.CollectVideo(userID, uint(videoID))
		if err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err1.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "收藏成功",
		})
	case "4": //转发
		if video_ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "videoID为空",
			})
			return
		}
		videoID, err := strconv.ParseUint(video_ID, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		video, err := service.SearchVideo(uint(videoID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		if video.InBlackList {
			c.JSON(http.StatusOK, gin.H{
				"msg": "视频未审核",
			})
			return
		}
		err1 := service.ForwardVideo(userID, uint(videoID))
		if err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err1.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "转发成功",
		})
	case "5": //弹幕
		if video_ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "videoID为空",
			})
			return
		}
		videoID, err := strconv.ParseUint(video_ID, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		video, err := service.SearchVideo(uint(videoID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err:": err.Error(),
			})
			return
		}
		if video.InBlackList {
			c.JSON(http.StatusOK, gin.H{
				"msg": "视频未审核",
			})
			return
		}
		if context == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "context为空",
			})
			return
		}
		err1 := service.CreateBarrage(userID, uint(videoID), context)
		if err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err1.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "发弹幕成功",
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "action错误",
		})
		return
	}
}
