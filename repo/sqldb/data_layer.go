package sqldb

import (
	"database/sql"

	"github.com/robertsmoto/db_controller_example/repo/models"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	DB *sql.DB
}

// NewUserRepo ..
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

// FindByID ..
func (r *UserRepo) FindByID(ID int) (*models.User, error) {
	return &models.User{}, nil
}

// Save ..
func (r *UserRepo) Save(user *models.User) error {
	return nil
}
