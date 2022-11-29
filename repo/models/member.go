package models

type Member struct {
	BaseData
	Type MemberType `json:"type" validate:"required"`
}

// MemberDataLayer interface ..
type MemberDataAccessLayer interface {
	Get(args []interface{}) (interface{}, error)
	Modify(Key string, args ...interface{}) (*Member, error)
	Save(member *Member) error
}
