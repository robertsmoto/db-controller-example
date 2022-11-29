package middleware

import (
	"github.com/gomodule/redigo/redis"
	"github.com/robertsmoto/db_controller_example/config"
)

type MiddlewareConn struct {
	DB   *redis.Conn
	conf *config.Config
}

func NewMiddlewareConnector(db *redis.Conn, conf *config.Config) *MiddlewareConn {
	mc := MiddlewareConn{
		DB:   db,
		conf: conf,
	}
	return &mc
}
