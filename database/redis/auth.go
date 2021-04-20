package database

import (
	"errors"
	"hackthon/database"
	"time"
)

var (
	NotFounded = errors.New("not founded")
)

// CreateTokenCache 创建token的缓存，通过该缓存来进行续签令牌操作
func CreateTokenCache(username, token string) error {
	conn, ctx := database.Redis()
	err := conn.Set(ctx, token, time.Now().Unix(), time.Hour*4).Err()
	if err != nil {
		return err
	}
	//设置此处的key是为了在修改密码后，能找到对应的token缓存并将其删除，然后把自身修改成新的token
	//如果在创建缓存的时候发现之前有token的缓存，就把缓存给删除了，然后更新这个username的缓存
	if ok := conn.Exists(ctx, username).Val(); ok == 1 {
		token, err := conn.Get(ctx, username).Result()
		if err != nil {
			return err
		}
		err = conn.Del(ctx, token).Err()
		if err != nil {
			return err
		}
	}
	err = conn.Set(ctx, username, token, time.Hour*4).Err()
	return err
}

func FindTokenCache(token string, time time.Duration) (time.Duration, error) {
	conn, ctx := database.Redis()
	if ok := conn.Exists(ctx, token).Val(); ok == 0 {
		return 0, NotFounded
	}
	ttl := conn.TTL(ctx, token).Val()
	return time - ttl, nil

}

func StoreCode(code string) error {
	conn, ctx := database.Redis()
	err := conn.Set(ctx, code, time.Now().Unix(), time.Minute*15).Err()
	return err
}

func FindCode(code string) error {
	conn, ctx := database.Redis()
	if ok := conn.Exists(ctx, code).Val(); ok == 0 {
		return NotFounded
	} else {
		return nil
	}
}

func DeleteCode(code string) error {
	conn, ctx := database.Redis()
	if ok := conn.Del(ctx, code).Val(); ok == 0 {
		return NotFounded
	} else {
		return nil
	}
}
