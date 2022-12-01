package dal

import (
	"github.com/gomodule/redigo/redis"
	"github.com/robertsmoto/db_controller_example/config"
	"github.com/robertsmoto/db_controller_example/repo/models"
)

// MemberDAL implements the models.MemberDataAccessLayer interface
type BaseDAL struct {
	DB   *redis.Conn
	conf *config.Config
}

func NewBaseDAL(db *redis.Conn, conf *config.Config) *BaseDAL {
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

/*
These may not be the best function signatures for a relational db, but
I'm assuming they could be made to work. You would most likely want to return
a model instance rather than in interface{}
*/

// Get ..
func (dal *MemberDAL) Get(args []interface{}) (interface{}, error) {
	// not implemented
	return nil, nil
}

// Modify ..
func (dal *MemberDAL) Modify(args []interface{}) error {
	return nil
}

// Save ..
func (dal *MemberDAL) Save(m *models.Member) error {
	return nil
}
