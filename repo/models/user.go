package models

// User ..
type User struct {
	Name string
}

// UserDataLayer interface ..
type UserDataLayer interface {
	GetByID(ID int) (*User, error)
	SaveJson(user *User) error
}
