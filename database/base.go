package database

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"hackthon/model"
	"log"
	"os"
	"time"
)

var db *gorm.DB
var rdb *redis.Client

func ConnectDB() {
	var err error
	var sqlDB *sql.DB
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	mysqlConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //建表命名约束为根据结构体名称单数小写命名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	}
	user := os.Getenv("MYSQL_USER")
	passwd := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: user + ":" + passwd + "@tcp(localhost:3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local",
		//DSN:                       "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameColumn:   true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}), &mysqlConfig)
	if err != nil {
		log.Println(err.Error())
		return
	}
	sqlDB, err = db.DB()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//设置连接池最大连接数
	sqlDB.SetMaxOpenConns(1000)
	//设置连接池最大闲置连接数
	sqlDB.SetMaxIdleConns(100)
	//设置连接最大闲置时间
	sqlDB.SetConnMaxIdleTime(time.Hour)
	//设置连接最大存活时间
	sqlDB.SetConnMaxLifetime(2 * time.Hour)
	err = db.AutoMigrate(&model.User{}, &model.Diary{}, &model.Todo{}, &model.Watch{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func DB() *gorm.DB {
	ctx := context.Background()
	return db.WithContext(ctx)
}

// DisconnectDB 断开与数据库的连接
func DisconnectDB() {
	mysqlDB, _ := db.DB()
	if err := mysqlDB.Close(); err != nil {
		log.Println(err.Error())
	}
}
func ConnectRedis() {
	passwd := os.Getenv("REDIS_PASSWD")
	rdb = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    passwd,
		DB:          0,
		IdleTimeout: -1,

		DialTimeout: time.Second * 5,
		PoolSize:    200,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err.Error())
	}
}
func Redis() (*redis.Conn, context.Context) {
	ctx := context.Background()
	return rdb.Conn(ctx), ctx
}
func DisconnectRedis() {
	if err := rdb.Close(); err != nil {
		log.Fatal(err.Error())
	}
}
