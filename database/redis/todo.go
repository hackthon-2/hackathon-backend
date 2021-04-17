package database

import (
	"hackthon/database"
	"strconv"
	"time"
)

// CreateTodoCache 接收序列化之后的数据进行存储
func CreateTodoCache(userId uint, data, date string) error {
	conn, ctx := database.Redis()
	defer conn.Close()
	err := conn.HSet(ctx, "USER_TODO_"+strconv.Itoa(int(userId)), date, data).Err()
	if err != nil {
		return err
	}
	return conn.Expire(ctx, "USER_TODO_"+strconv.Itoa(int(userId)), time.Hour*72).Err()
}
func DeleteTodoFieldCache(userId uint, date string) error {
	conn, ctx := database.Redis()
	defer conn.Close()
	err := conn.HDel(ctx, "USER_TODO_"+strconv.Itoa(int(userId)), date).Err()
	if err != nil {
		return err
	}
	return conn.Expire(ctx, "USER_TODO_"+strconv.Itoa(int(userId)), time.Hour*72).Err()
}

// FindTodoCache 查找缓存
func FindTodoCache(userId uint, date string) (string, error) {
	conn, ctx := database.Redis()
	defer conn.Close()
	data, err := conn.HGet(ctx, "USER_TODO_"+strconv.Itoa(int(userId)), date).Result()
	if err != nil {
		return "", err
	}
	conn.Expire(ctx, "USER_TODO_"+strconv.Itoa(int(userId)), time.Hour*72)
	return data, nil
}

// CreateStatisticsCache 建立统计的缓存
func CreateStatisticsCache(userId uint, from, data string) error {
	conn, ctx := database.Redis()
	err := conn.Set(ctx, "USER_"+strconv.Itoa(int(userId))+"_STAT_FROM_"+from, data, time.Hour*24).Err()
	return err
}

// FindStatisticsCache 查找统计的缓存
func FindStatisticsCache(userId uint, from string) (string, error) {
	conn, ctx := database.Redis()
	defer conn.Close()
	data, err := conn.Get(ctx, "USER_"+strconv.Itoa(int(userId))+"_STAT_FROM_"+from).Result()
	if err != nil {
		return "", err
	}
	return data, nil
}
