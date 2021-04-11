package database

import (
	"errors"
	"hackthon/database"
	"time"
)
var NotFounded=errors.New("not founded")

func CreateTokenCache(token string) error {
	conn, ctx := database.Redis()
	err := conn.HSet(ctx, token, "lock", 0).Err()
	if err != nil {
		return err
	}
	err = conn.Expire(ctx, token, time.Hour*4).Err()
	defer conn.Close()
	return err
}

func FindTokenCache(token string) (string, error) {
	conn, ctx := database.Redis()
	if ok:=conn.Exists(ctx, token).Val();ok==0{
		return "",NotFounded
	}
	lock, err := conn.HGetAll(ctx, token).Result()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return lock["lock"], nil
}
