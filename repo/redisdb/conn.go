package redisdb

import (
	"github.com/gomodule/redigo/redis"
)

// DB is a global variable to hold db connection
var DB *redis.Conn

type DBInstance int

const (
	Auth DBInstance = iota
	Api
	Data
)

// ConnectDB opens a connection to the database
func ConnectDB(dbInstance DBInstance) *redis.Conn {

	netw := ""
	addr := ""

	conn, err := redis.Dial(
		netw,
		addr,
		redis.DialDatabase(int(dbInstance)),
		redis.DialPassword(conf),
	)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
