package repositories

import (
	"database/sql"

	"github.com/robertsmoto/db_controller_example/models"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo ..
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
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
