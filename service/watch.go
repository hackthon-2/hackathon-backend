package service

import (
	"errors"
	database "hackthon/database/mysql"
	database2 "hackthon/database/redis"
	"hackthon/model"
	"hackthon/util"
)

var (
	ErrorWhenCreatingWatch = errors.New("when creating watch had a error")
	ErrorWhenUpdatingWatch = errors.New("when updating watch had a error")
	FindWatchError         = errors.New("found watch error")
	ErrorWhenDeletingWatch = errors.New("when deleting watch had a error")
)

func CreateWatch(username string, userId uint, input *model.WatchInput) error {
	var watch model.Watch
	util.StructAssign(&watch, input)
	watch.Username = username
	watch.UserID = userId
	row, err := database.CreateWatch(&watch)
	if err != nil {
		return err
	}
	if row != 1 {
		return ErrorWhenCreatingWatch
	}
	data, row, err := database.FindWatch(userId)
	if err != nil {
		return err
	}
	if row != 1 {
		return ErrorWhenCreatingWatch
	}
	err = database2.CreateWatchCache(userId, &data)
	return err
}

func UpdateWatch(watchId, userId uint, input *model.UpdateWatchInput) error {
	var watch model.Watch
	util.StructAssign(&watch, input)
	row, err := database.UpdateWatch(watchId, &watch)
	if err != nil {
		return err
	}
	if row != 1 {
		return ErrorWhenUpdatingWatch
	}
	data, row, err := database.FindWatch(userId)
	if err != nil {
		return err
	}
	if row != 1 {
		return ErrorWhenUpdatingWatch
	}
	err = database2.CreateWatchCache(userId, &data)
	return err
}

func FindWatch(userId uint) (model.Watch, error) {
	data, err := database2.FindWatchCache(userId)
	if err != nil {
		dat, row, e := database.FindWatch(userId)
		if row != 1 {
			return model.Watch{}, FindWatchError
		}
		if e != nil {
			return model.Watch{}, e
		}
		err = database2.CreateWatchCache(userId, &dat)
		if err != nil {
			return model.Watch{}, err
		}
		return dat, nil
	}
	return data, nil
}

func DeleteWatch(userId uint) error {
	data, row, err := database.FindWatch(userId)
	if err != nil {
		return ErrorWhenDeletingWatch
	}
	if row != 1 {
		return ErrorWhenDeletingWatch
	}
	err = database.DeleteWatch(data.ID)
	if err != nil {
		return err
	}
	err = database2.DeleteWatchCache(userId)
	return err
}
