package database

import (
	"hackthon/database"
	"hackthon/model"
)

// CreateUser 创建用户数据
func CreateUser(user *model.User) (error, int64) {
	result := database.DB().Table("user").Create(user)
	return result.Error, result.RowsAffected
}

// FindUserById 根据用户id查找用户信息
func FindUserById(id uint) (model.User, int64, error) {
	var user model.User
	result := database.DB().Table("user").Select("id","username","email","avatar").First(&user, id)
	return user, result.RowsAffected, result.Error
}

// FindUserByUsername 根据用户名查找用户信息
func FindUserByUsername(username string) (model.User, int64, error) {
	var user model.User
	result := database.DB().Table("user").Where("username=?", username).First(&user)
	return user, result.RowsAffected, result.Error
}

// FindUserByEmail 根据邮箱地址查找用户信息
func FindUserByEmail(email string) ([]model.User, error) {
	var user []model.User
	result := database.DB().Table("user").Where("email=?", email).Find(&user)
	return user, result.Error
}

// UpdateUserById 更新用户信息，比如密码和头像，还有用户名,电子邮箱
func UpdateUserById(id uint, user *model.User) (int64, error) {
	result := database.DB().Table("user").Where(model.User{ID: id}).Updates(user)
	return result.RowsAffected, result.Error
}
