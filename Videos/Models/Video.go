package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title       string `gorm:"not null;"`
	URL         string `gorm:"not null;"`
	UserID      uint
	InBlackList bool
	ThumbsUpNum uint
	IsOriginal  bool
	OriUserId   uint
	Type        string `gorm:"not null;"`
}

type Comment struct {
	gorm.Model
	UserID       uint
	VideoID      uint
	ParentID     uint
	RootParentID uint
	Context      string     `gorm:"not null;"`
	Parent       *Comment   `gorm:"foreignkey:ParentID"`
	Children     []*Comment `gorm:"foreignkey:ParentID"`
}

type Barrage struct {
	gorm.Model
	UserID  uint
	VideoID uint
	Context string `gorm:"not null;"`
}

type Favorite struct {
	UserID  uint
	VideoID uint
}

//自动生成视频id

// 上传视频
func UploadVideo(userID uint, URL string, Type string, title string) *gorm.DB {
	video := Video{UserID: userID, URL: URL, InBlackList: true, Type: Type, ThumbsUpNum: 0, IsOriginal: true, Title: title}
	return db.Create(&video)
}

// 视频过审
func ReleaseVideo(videoID uint) *gorm.DB {
	return db.Model(&Video{}).Where("id=?", videoID).Update("InBlackList", false)
}

// 查询已审核视频
func SearchPassVideo(videoID uint) (Video, *gorm.DB) {
	video := Video{}
	db := db.Model(&Video{}).Where("id=? AND in_black_list=?", videoID, false).Find(&video)
	return video, db
}

// 查询视频
func SearchVideo(videoID uint) (Video, *gorm.DB) {
	video := Video{}
	db := db.Model(&Video{}).Where("id=?", videoID).Find(&video)
	return video, db
}

// 按类别查询视频
func SearchVideoType(Type string) ([]Video, *gorm.DB) {
	var videos []Video
	db := db.Model(&Video{}).Where("Type=? AND in_black_list=?", Type, false).Find(&videos)
	return videos, db
}

// 按发布时间查询视频
func SearchVideoMandD(month time.Month, day int) ([]Video, *gorm.DB) {
	var videos []Video
	db := db.Model(&Video{}).Where("MONTH(created_at) =? AND DAY(created_at)=? AND in_black_list=?", month, day, false).Find(&videos)
	return videos, db
}

// 按年份查询视频
func SearchVideoYear(year int) ([]Video, *gorm.DB) {
	var videos []Video
	db := db.Model(&Video{}).Where("year(created_at) = ? AND in_black_list=?", year, false).Find(&videos)
	return videos, db
}

// 排序查询视频
// 点击量
func SearchVideoThumbupOrd(title string) ([]Video, *gorm.DB) {
	var videos []Video
	db := db.Model(&Video{}).Where("title LIKE ? AND in_black_list=?", "%"+title+"%", false).Order("thumbs_up_num desc").Find(&videos)
	return videos, db
}

// 发布时间
func SearchVideoTimeOrd(title string) ([]Video, *gorm.DB) {
	var videos []Video
	db := db.Model(&Video{}).Where("title LIKE ? AND in_black_list=?", "%"+title+"%", false).Order("create_at desc").Find(&videos)
	return videos, db
}

// 转发视频
// 复制视频相关信息
func CopyVideo(video Video) (Video, *gorm.DB) {
	var video1 = Video{URL: video.URL, InBlackList: false, ThumbsUpNum: 0, IsOriginal: false, OriUserId: video.UserID}
	db := db.Create(&video1)
	return video1, db
}

// 修改是否原创信息
func ChangeIsOri(videoID uint) *gorm.DB {
	return db.Model(&Video{}).Where("id=?", videoID).Update("IsOriginal", true)

}

// 修改原创者信息
func ChangeOriUser(videoID uint, oriUserID uint) *gorm.DB {
	return db.Model(&Video{}).Where("id=?", videoID).Update("OriUserID", oriUserID)
}

// 修改视频发布人
func ChangeVideoUserID(userID uint, videoID uint) *gorm.DB {
	return db.Model(&Video{}).Where("id=?", videoID).Update("UserID", userID)
}

// 点赞视频
func ThumbUpVideo(videoID uint) *gorm.DB {
	return db.Model(&Video{}).Where("id=? AND in_black_list=?", videoID, false).Update("thumbs_up_num", gorm.Expr("thumbs_up_num+ ?", 1))
}

// 评论
// 创建评论
func CommentVideo(userID uint, context string) (Comment, *gorm.DB) {
	comment := Comment{UserID: userID, Context: context}
	db := db.Create(&comment)
	return comment, db
}

// 查询评论
func SearchComment(commentID uint) (Comment, *gorm.DB) {
	comment := Comment{}
	res := db.First(&comment, commentID)
	return comment, res
}

// 按父节点查询评论
func SearchCommentByParID(parentID uint) ([]Comment, *gorm.DB) {
	var comments []Comment
	db := db.Model(&Comment{}).Where("parent_id=?", parentID).Find(&comments)
	return comments, db
}

// 修改评论的videoID
func ChangeVideoID(commentID uint, videoID uint) *gorm.DB {
	return db.Model(&Comment{}).Where("id=?", commentID).Update("VideoID", videoID)
}

// 删除根节点为所需的所有评论
func DeleteRootParentCom(rootParentID uint) *gorm.DB {
	return db.Where("root_parent_id=?", rootParentID).Delete(&Comment{})
}

//按userID批量删除

// 修改父评论
func ChangeParentID(commentID uint, parentID uint) *gorm.DB {
	return db.Model(&Comment{}).Where("id=?", commentID).Update("ParentID", parentID)
}

// 修改根评论
func ChangeRootParentID(commentID uint, rootParentID uint) *gorm.DB {
	return db.Model(&Comment{}).Where("id=?", commentID).Update("RootParentID", rootParentID)
}

// 创建弹幕
func BarrageVideo(userID uint, videoID uint, context string) *gorm.DB {
	barrage := Barrage{UserID: userID, VideoID: videoID, Context: context}
	return db.Create(&barrage)
}

// 收藏视频
func Collect(userID uint, videoID uint) *gorm.DB {
	favorite := Favorite{UserID: userID, VideoID: videoID}
	return db.Create(&favorite)
}

// 根据userID批量删除
// 视频
func DeleteVideoByUserID(userID uint) *gorm.DB {
	return db.Model(&Video{}).Where("user_id=?", userID).Delete(&Video{})
}

// 评论
func DeleteCommentByUserID(userID uint) *gorm.DB {
	return db.Model(&Comment{}).Where("user_id=?", userID).Delete(&Comment{})
}

// 弹幕
func DeleteBarrageByUserID(userID uint) *gorm.DB {
	return db.Model(&Barrage{}).Where("user_id=?", userID).Delete(&Barrage{})
}

// 收藏
func DeleteFavoriteByUserID(userID uint) *gorm.DB {
	return db.Where("user_id=?", userID).Delete(&Favorite{})
}

// 查询所有子评论ID
func SearchAllChildrenID(parentID uint) ([]uint, *gorm.DB) {
	var childrenID []uint
	db := db.Model(&Comment{}).Where("parent_id = ?", parentID).Pluck("id", &childrenID)
	return childrenID, db
}

// 删除评论
func DeleteComment(commentID uint) *gorm.DB {
	return db.Delete(&Comment{}, commentID)
}
