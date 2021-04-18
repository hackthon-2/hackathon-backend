package database

import (
	"hackthon/database"
	"hackthon/model"
)

// CreateWatch 创建watch表格
func CreateWatch(watch *model.Watch) (int64, error) {
	result := database.DB().Create(watch)
	return result.RowsAffected, result.Error
}

// UpdateWatch 更新watch表格
func UpdateWatch(watchId uint, watch *model.Watch) (int64, error) {
	result := database.DB().Table("watch").Where("id=?", watchId).Updates(watch)
	return result.RowsAffected, result.Error
}

// FindWatch 查找当前用户是否有存在的表项
func FindWatch(userId uint) (model.Watch, int64, error) {
	var watch model.Watch
	result := database.DB().Table("watch").Where("user_id=?", userId).First(&watch)
	return watch, result.RowsAffected, result.Error
}

// ListWatchByTime 根据建立的监督所需要完成的次数来匹配用户
func ListWatchByTime(times uint) ([]model.Watch, error) {
	var watches []model.Watch
	result := database.DB().Table("watch").Where("time=?", times).Limit(5).Find(&watches)
	return watches, result.Error
}

// DeleteWatch 删除监督条项
func DeleteWatch(watchId uint) error {
	result := database.DB().Table("watch").Where("id=?", watchId).Delete(&model.Watch{})
	return result.Error
}

