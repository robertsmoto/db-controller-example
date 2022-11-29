package redisdb

import (
	"fmt"

	"github.com/gomodule/redigo/redis"

	"github.com/robertsmoto/db_controller_example/config"
	"github.com/robertsmoto/db_controller_example/repo/models"
)

// MemberDAL implements the models.MemberDataAccessLayer interface
type BaseDAL struct {
	DB   redis.Conn
	conf *config.Config
}

func NewBaseDAL(db redis.Conn, conf *config.Config) *BaseDAL {
	return &BaseDAL{
		DB:   db,
		conf: conf,
	}
}

type MemberDAL struct {
	BaseDAL
}

// Contructor
func NewMemberDAL(dal *BaseDAL) *MemberDAL {
	return &MemberDAL{
		BaseDAL: *dal,
	}
}

// Get ..
func (dal *MemberDAL) Get(args []interface{}) (interface{}, error) {
	fmt.Printf("\n## paths --> %[1]v %[1]T", args)
	// add ID to various collections, and possibly organize search here
	return dal.DB.Do("JSON.GET", args...)
}

// Modify ..
func (dal *MemberDAL) Modify(args []interface{}) error {
	return nil
}

// Save ..
func (dal *MemberDAL) Save(member *models.Member) error {
	return nil
}
