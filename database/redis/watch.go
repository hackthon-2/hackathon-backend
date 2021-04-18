package database

import (
	"encoding/json"
	"hackthon/database"
	"hackthon/model"
	"strconv"
	"time"
)

func CreateWatchCache(userId uint, watch *model.Watch) error {
	val, err := json.Marshal(watch)
	if err != nil {
		return err
	}
	data := string(val)
	conn, ctx := database.Redis()
	return conn.Set(ctx, "USER_WATCH_"+strconv.Itoa(int(userId)), data, time.Hour*72).Err()
}

func FindWatchCache(userId uint) (model.Watch, error) {
	conn, ctx := database.Redis()
	val, err := conn.Get(ctx, "USER_WATCH_"+strconv.Itoa(int(userId))).Result()
	if err != nil {
		return model.Watch{}, err
	}
	var watch model.Watch
	err = json.Unmarshal([]byte(val), &watch)
	return watch, err
}

func DeleteWatchCache(userId uint) error {
	conn, ctx := database.Redis()
	return conn.Del(ctx, "USER_WATCH_"+strconv.Itoa(int(userId))).Err()
}
