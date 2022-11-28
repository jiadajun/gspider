package common

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

var LocalSpider = &DbStruct{
	User:           "root",
	Password:       "123",
	Host:           "127.0.0.1",
	Port:           3306,
	Database:       "spider",
	ConnectTimeout: "10s",
	Charset:        "utf8",
}

var AtSpider = &DbStruct{
	User:           "at",
	Password:       "at111",
	Host:           "192.168.2.114",
	Port:           3306,
	Database:       "1233",
	ConnectTimeout: "10s",
	Charset:        "utf8",
}

var RedisClient1 = redis.NewClient(&redis.Options{Addr: "111111111118:6379", Password: "111111111111"})
