package database

import (
	"encoding/json"
	"errors"
	"hackthon/database"
	"hackthon/model"
	"strconv"
	"time"
)

var NullPointer = errors.New("null pointer")

// CreateDiaryCache 刷新缓存机制
func CreateDiaryCache(diary *[]model.Diary) error {
	if diary == nil {
		return NullPointer
	}
	val, err := json.Marshal(diary)
	if err != nil {
		return err
	}
	conn, ctx := database.Redis()
	err = conn.HSet(ctx, "USER_"+strconv.Itoa(int((*diary)[0].UserID)), (*diary)[0].Time, string(val)).Err()
	if err != nil {
		return err
	}
	return conn.Expire(ctx, "USER_"+strconv.Itoa(int((*diary)[0].UserID)), time.Hour*72).Err()
}

// FindDiaryCache 查找缓存，没有就去数据库找，找了之后创建缓存
func FindDiaryCache(userId uint, date string) (string, error) {
	conn, ctx := database.Redis()
	return conn.HGet(ctx, "USER_"+strconv.Itoa(int(userId)), date).Result()
}
