package service

import (
	"encoding/json"
	"errors"
	database "hackthon/database/mysql"
	database2 "hackthon/database/redis"
	"hackthon/model"
	"hackthon/util"
)

var (
	ErrorWhenCreateDiary = errors.New("when creating diary had a error")
	ErrorWhenUpdateDiary = errors.New("when creating diary had a error")
	GetDiaryListError    = errors.New("getting diary list failed")
	GetDiaryFailed       = errors.New("getting diary failed")
	DeleteDiaryError     = errors.New("can't delete diary")
)

// CreateDiary 创建日记数据
func CreateDiary(userID uint, input *model.DiaryInput) error {
	var diary model.Diary
	util.StructAssign(&diary, input)
	diary.UserID = userID
	err, row := database.CreateDiary(&diary)
	if err != nil {
		return err
	}
	if row != 1 {
		return ErrorWhenCreateDiary
	}
	//创建了一项日记后，把今天所有的日记项全部拿出来，建立缓存&更新缓存
	data, err := database.ListDiaryByTime(userID, input.Time)
	if err != nil {
		return err
	}
	err = database2.CreateDiaryCache(&data)
	return err
}

// UpdateDiary 更新日记数据
func UpdateDiary(userID, diaryID uint, input *model.DiaryInput) error {
	var diary model.Diary
	util.StructAssign(&diary, input)
	//更新数据库记录，通过id来查找要更新的记录
	err, row := database.UpdateDiary(diaryID, &diary)
	if err != nil {
		return err
	}
	if row != 1 {
		return ErrorWhenUpdateDiary
	}
	//更新完了就更新当天的缓存
	data, err := database.ListDiaryByTime(userID, input.Time)
	if err != nil {
		return err
	}
	err = database2.CreateDiaryCache(&data)
	return err
}

func FindDiary(diaryID uint) (model.Diary, error) {
	diary, row, err := database.FindDiaryById(diaryID)
	if err != nil {
		return model.Diary{}, err
	}
	if row != 1 {
		return model.Diary{}, GetDiaryFailed
	}
	return diary, nil
}

// ListDiary 列出日记数据
func ListDiary(userID uint, date string) ([]model.Diary, error) {
	data, err := database2.FindDiaryCache(userID, date)
	if err != nil || data == "" {
		//没找到缓存就到数据库拿
		dat, e := database.ListDiaryByTime(userID, date)
		if e != nil {
			return []model.Diary{}, GetDiaryListError
		}
		if len(dat) < 1 {
			return []model.Diary{}, nil
		}
		e = database2.CreateDiaryCache(&dat)
		//创建缓存失败直接返回错误
		if e != nil {
			return nil, e
		}
		return dat, nil
	}
	if data == "" {
		return []model.Diary{}, nil
	}
	var diary []model.Diary
	//把提取到的缓存字符串（json格式)反序列化成切片
	_ = json.Unmarshal([]byte(data), &diary)
	return diary, nil
}

// DeleteDiary 删除日记数据
func DeleteDiary(userID, diaryID uint) error {
	//先看看要删除的数据是否存在
	diary, row, err := database.FindDiaryById(diaryID)
	if row != 1 {
		return DeleteDiaryError
	}
	if err != nil {
		return err
	}
	//存在就删除
	err = database.DeleteDiary(diaryID)
	if err != nil {
		return err
	}
	//然后找到当天的记录
	data, err := database.ListDiaryByTime(userID, diary.Time)
	if err != nil {
		return err
	}
	if len(data) < 1 {
		return database2.DeleteFieldDiaryCache(userID, diary.Time)
	}
	//建立更新后的缓存
	return database2.CreateDiaryCache(&data)
}
