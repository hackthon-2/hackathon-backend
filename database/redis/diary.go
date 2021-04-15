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
	//打平diary切片成json
	val, err := json.Marshal(diary)
	if err != nil {
		return err
	}
	conn, ctx := database.Redis()
	defer conn.Close()
	//设置缓存，正常返回码是如果建立新的则返回1，替换旧的则返回0，所以这里只看有没有错误
	err = conn.HSet(ctx, "USER_"+strconv.Itoa(int((*diary)[0].UserID)), (*diary)[0].Time, string(val)).Err()
	if err != nil {
		return err
	}
	//设置用户的缓存，保存三天，如果三天都没有对缓存进行任何操作则自动失效
	return conn.Expire(ctx, "USER_"+strconv.Itoa(int((*diary)[0].UserID)), time.Hour*72).Err()
}

// FindDiaryCache 查找缓存，缓存是一个hashTable,可以根据日期存缓存字段
func FindDiaryCache(userId uint, date string) (string, error) {
	conn, ctx := database.Redis()
	defer conn.Close()
	//先找缓存是否存在
	data, err := conn.HGet(ctx, "USER_"+strconv.Itoa(int(userId)), date).Result()
	//如果出现错误就直接返回错误
	if err != nil {
		return "", err
	}
	//如果找到了，则将缓存有效期更新
	conn.Expire(ctx, "USER_"+strconv.Itoa(int(userId)), time.Hour*72)
	return data, nil
}
