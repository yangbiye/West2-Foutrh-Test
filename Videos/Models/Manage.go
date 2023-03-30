package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Password string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(50);not null"`
}

// 管理员登录之后给出信息
func LoginManager(managerID uint) (Manager, *gorm.DB) {
	manager := Manager{}
	result := db.First(&manager, managerID)
	return manager, result
}

// 管理员注册
func RegisterManager(password string, email string) (Manager, *gorm.DB) {
	manager := Manager{Password: password, Email: email}
	db := db.Create(&manager)
	return manager, db
}
