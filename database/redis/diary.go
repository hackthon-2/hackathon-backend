package database

import (
	"encoding/json"
	"hackthon/database"
	"hackthon/model"
	"strconv"
)

// CreateDiaryCache 刷新缓存机制
func CreateDiaryCache(diary *[]model.Diary) error {
	val, err := json.Marshal(diary)
	if err != nil {
		return err
	}
	conn, ctx := database.Redis()
	return conn.HSet(ctx, "USER_"+strconv.Itoa(int((*diary)[0].UserID)), (*diary)[0].Time, string(val)).Err()
}
