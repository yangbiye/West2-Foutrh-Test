package service

import (
	"time"

	models "github.com/videos/Models"
)

// 搜索视频
// 按年份搜索视频
func SearchVideoYear(year int) ([]models.Video, error) {
	videos, res := models.SearchVideoYear(year)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return videos, nil
	}
}

// 按类别搜索
func SearchVideoType(Type string) ([]models.Video, error) {
	videos, res := models.SearchVideoType(Type)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return videos, nil
	}
}

// 按发布时间搜索
func SearchVideoMandD(month time.Month, day int) ([]models.Video, error) {
	videos, res := models.SearchVideoMandD(month, day)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return videos, nil
	}
}

// 搜索用户
// 按地点搜索
func SearchUserPlace(place string) ([]models.User, error) {
	users, res := models.SearchUserPlace(place)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return users, nil
	}
}

// 查看历史搜索记录
func SearchRecord() ([]models.Record, error) {
	records, res := models.SearchRecord()
	if res.Error != nil {
		return nil, res.Error
	} else {
		return records, nil
	}
}

// 筛选视频
// 点赞量
func SearchVideoThumbupOrd(title string) ([]models.Video, error) {
	videos, res := models.SearchVideoThumbupOrd(title)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return videos, nil
	}
}

// 发布时间先后
func SearchVideoTimeOrd(title string) ([]models.Video, error) {
	videos, res := models.SearchVideoTimeOrd(title)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return videos, nil
	}
}
