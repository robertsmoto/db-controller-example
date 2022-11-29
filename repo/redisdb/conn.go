package redisdb

import (
	"github.com/gomodule/redigo/redis"
	"github.com/robertsmoto/db_controller_example/config"
)

// DB is a global variable to hold db connection
var DB *redis.Conn

type DBInstance int

const (
	Account DBInstance = iota + 5
	Api
	Data
)

// ConnectDB opens a connection to the database
func ConnectDB(dbInstance DBInstance) redis.Conn {
	conf := config.Conf

	conn, err := redis.Dial(
		conf.RedisCredentials.Netw,
		conf.RedisCredentials.Addr,
		redis.DialDatabase(int(dbInstance)),
		redis.DialPassword(conf.RedisCredentials.Pass),
	)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
