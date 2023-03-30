package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password     string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"type:varchar(256);not null"`
	UserName     string `gorm:"type:varchar(256);"`
	Introduction string `gorm:"type:varchar(255)"`
	Place        string `gorm:"type:varchar(255)"`
	Picture      string `gorm:"type:varchar(255)"`
	InBlackList  bool   `gorm:"not null"`
}

type Record struct {
	gorm.Model
	Rec string `gorm:"type:varchar(255);not null"`
}

// 用户登录之后给出用户信息
func LoginUser(userID uint) (User, *gorm.DB) {
	user := User{}
	result := db.First(&user, userID)
	return user, result
}

// 用户注册
func RegisterUser(password string, email string) (User, *gorm.DB) {
	user := User{Password: password, Email: email, InBlackList: false, UserName: "", Introduction: "", Place: "", Picture: ""}

	result := db.Create(&user)

	return user, result
}

// 个人资料

// 更新用户名
func SetUserName(userID uint, userName string) *gorm.DB {
	return db.Model(&User{}).Where("id=?", userID).Update("UserName", userName)
}

// 更新简介
func SetIntroduction(userID uint, introduction string) *gorm.DB {
	return db.Model(&User{}).Where("id=?", userID).Update("Introduction", introduction)

}

// 更新生日
func SetBirthday(userID uint, time *time.Time) *gorm.DB {

	return db.Model(&User{}).Where("id=?", userID).Update("Birthhday", time)

}

// 更新所在地
func SetPlace(userID uint, place string) *gorm.DB {
	return db.Model(&User{}).Where("id=?", userID).Update("Place", place)

}

// 修改密码
func SetPassword(userID uint, password string) *gorm.DB {
	return db.Model(&User{}).Where("id=?", userID).Update("Password", password)

}

// 更新头像
func SetPicture(userID uint, picture string) *gorm.DB {
	return db.Model(&User{}).Where("id=?", userID).Update("Picture", picture)

}

// 邮箱
// 用户拉黑
func LockUser(userID uint) *gorm.DB {
	return db.Model(&User{}).Where("id=?", userID).Update("InBlackList", true)
}

// 拉白用户
func DeBockUser(userID uint) *gorm.DB {
	return db.Model(&User{}).Where("id=?", userID).Update("InBlackList", false)
}

// 软删除用户(封号)
func DeleteUser(userID uint) *gorm.DB {
	return db.Delete(&User{}, userID)
}

// 查找软删除用户
func FindBlackUser(userID uint) (User, *gorm.DB) {
	user := User{}
	res := db.Unscoped().Where("UserID = ?", userID).Find(&user)
	return user, res
}

// 按地点查找用户
func SearchUserPlace(place string) ([]User, *gorm.DB) {
	var users []User
	db := db.Model(&User{}).Where("Place=?", place).Find(&users)
	return users, db
}

// 查找用户
func SearchUSer(userID uint) (User, *gorm.DB) {
	var user User
	db := db.Model(&User{}).Where("id=?", userID).Find(&user)
	return user, db
}

// 创建搜索记录
func CreateRecord(rec string) *gorm.DB {
	record := Record{Rec: rec}
	return db.Create(&record)
}

// 查询所有历史记录
func SearchRecord() ([]Record, *gorm.DB) {
	var records []Record
	db := db.Order("id desc").Find(&records)
	return records, db
}
