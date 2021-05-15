package db

import (
	"errors"
	"pishikot/pkg/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	result := r.db.Find(&user, "login=?", user.Login)
	if result.RowsAffected > 0 {
		return errors.New("User already registered")
	}
	result = r.db.Create(&user)
	return nil
}

func (r *UserRepo) FindUser(user *models.User) error {
	result := r.db.First(&user, "login=? and password=?", user.Login, user.Password)
	if result.RowsAffected > 0 {
		return nil
	}
	return errors.New("User not found")
}

func (r *UserRepo) DeleteUser(user *models.User) error {
	result := r.db.Delete(user, "id=?", user.ID)
	if result.RowsAffected > 0 {
		return nil
	}
	return errors.New("User not exists")
}

func (r *UserRepo) GetUser(user *models.User) bool {
	result := r.db.First(&user, "id = ?", user.ID)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func (r *UserRepo) Clean(user *models.User) {
	user.Password = ""
}
