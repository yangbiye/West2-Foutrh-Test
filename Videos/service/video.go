package service

import models "github.com/videos/Models"

//上传视频
func UploadVideo(userID uint, URL string, Type string, title string) error {
	res := models.UploadVideo(userID, URL, Type, title)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

//点赞视频
func ThumbsUpVideo(videoID uint) error {
	res := models.ThumbUpVideo(videoID)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

//评论视频
func CommentVideo(userID uint, videoID uint, context string) error {
	comment1, res1 := models.CommentVideo(userID, context) //创建评论
	if res1.Error != nil {
		return res1.Error
	} else {
		res2 := models.ChangeVideoID(comment1.ID, videoID) //修改评论的videoID
		if res2.Error != nil {
			return res2.Error
		} else {
			return nil
		}
	}
}

//回复评论
func CommentComment(userID uint, commentID uint, context string) error {
	comment1, res1 := models.CommentVideo(userID, context) //创建评论
	if res1.Error != nil {
		return res1.Error
	} else {
		comment2, res2 := models.SearchComment(commentID) //查找父评论
		if res2.Error != nil {
			return res2.Error
		} else {
			res3 := models.ChangeParentID(comment1.ID, comment2.ID) //更改父评论ID
			if res3.Error != nil {
				return res3.Error
			} else {
				//更改根评论ID
				if comment2.RootParentID == 0 {
					res4 := models.ChangeRootParentID(comment1.ID, comment2.ID)
					if res4.Error != nil {
						return res4.Error
					}
				} else {
					res5 := models.ChangeRootParentID(comment1.ID, comment2.RootParentID)
					if res5.Error != nil {
						return res5.Error
					}
				}
				//更改评论的视频ID
				res6 := models.ChangeVideoID(comment1.ID, comment2.VideoID)
				if res6.Error != nil {
					return res6.Error
				} else {
					return nil
				}
			}
		}
	}
}

//收藏视频
func CollectVideo(userID uint, videoID uint) error {
	res := models.Collect(userID, videoID)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

//转发视频
func ForwardVideo(userID uint, videoID uint) error {
	video1, res1 := models.SearchPassVideo(videoID)
	if res1.Error != nil {
		return res1.Error
	} else {
		video2, res2 := models.CopyVideo(video1)
		if res2.Error != nil {
			return res2.Error
		} else {
			res3 := models.ChangeVideoUserID(userID, video2.ID)
			if res3.Error != nil {
				return res3.Error
			} else {
				return nil
			}
		}
	}
}

//发送弹幕
func CreateBarrage(userID uint, videoID uint, context string) error {
	res := models.BarrageVideo(userID, videoID, context)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

func DeleteVideoByUserID(UserID uint) error {
	res := models.DeleteVideoByUserID(UserID)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}
func DeleteCommentByUserID(UserID uint) error {
	res := models.DeleteCommentByUserID(UserID)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}
func DeleteBarrageByUserID(UserID uint) error {
	res := models.DeleteBarrageByUserID(UserID)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}
func DeleteFavoriteByUserID(UserID uint) error {
	res := models.DeleteFavoriteByUserID(UserID)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

//简单查询视频
func SearchVideo(videoID uint) (models.Video, error) {
	var video1 models.Video
	video, res := models.SearchVideo(videoID)
	if res.Error != nil {
		return video1, res.Error
	} else {
		return video, nil
	}
}

//俺父节点查询评论
// func SearchCommentByParID(parentID uint) ([]models.Comment, error) {
// 	comments, res := models.SearchCommentByParID(parentID)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	} else {
// 		return comments, nil
// 	}
// }
