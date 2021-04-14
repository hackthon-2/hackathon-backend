package database

import (
	"hackthon/database"
	"hackthon/model"
)

// CreateDiary 创建日记数据
func CreateDiary(diary *model.Diary) (error, int64) {
	result := database.DB().Table("diary").Create(diary)
	return result.Error, result.RowsAffected
}

// UpdateDiary 拉取列表后，会有id项，然后通过id项构造进去的链接，然后通过id项查找该项数据中的所有数据，从而对数据能够进行更新
func UpdateDiary(id uint, diary *model.Diary) (error, int64) {
	result := database.DB().Table("diary").Where("id=?", id).Updates(diary)
	return result.Error, result.RowsAffected
}

//FindDiaryById 根据id查找数据，主键索引，速度较快
func FindDiaryById(id uint) (model.Diary, int64, error) {
	var diary model.Diary
	result := database.DB().Table("diary").Where("id=?", id).First(&diary)
	return diary, result.RowsAffected, result.Error
}

// FindDiaryByUserIDAndQuestionAndTime 根据用户名，问题和日期查找日记(不一定用得上)
func FindDiaryByUserIDAndQuestionAndTime(userID uint, question, date string) (model.Diary, int64, error) {
	var diary model.Diary
	result := database.DB().Table("diary").Where("user_id=? AND time=? AND question=?", userID, date, question).First(&diary)
	return diary, result.RowsAffected, result.Error
}

// ListDiaryByTime 根据时间和userId查询diary
func ListDiaryByTime(userId uint, time string) ([]model.Diary, error) {
	var diary []model.Diary
	result := database.DB().Table("diary").Where("user_id=? AND time=?", userId, time).Find(&diary)
	if result.Error != nil {
		return nil, result.Error
	}
	//把日期的序列化格式改成YYYY-MM-DD
	for i, v := range diary {
		diary[i].Time = v.Time[:10]
	}
	return diary, result.Error
}

func DeleteDiary(diaryID uint) error {
	return database.DB().Delete(&model.Diary{}, diaryID).Error
}
