package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Login     string         `json:"login,omitempty" form:"login"`
	Password  string         `json:"password,omitempty" form:"password"`
	Email     string         `json:"email,omitempty" form:"email"`
	Age       int            `json:"age,omitempty" form:"age"`
	FirstName string         `json:"first_name,omitempty" form:"first_name"`
	LastName  string         `json:"last_name,omitempty" form:"last_name"`
	UserType  string         `json:"user_type,omitempty" form:"user_type"`
	LastLogin time.Time      `sql:"DEFAULT:current_timestamp"`
}

func (u User) ValidateForRegister() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Age, validation.Required, validation.Min(6)),
		validation.Field(&u.FirstName, validation.Required, validation.Length(4, 20)),
		validation.Field(&u.LastName, validation.Required, validation.Length(4, 20)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Login, validation.Required, validation.Length(4, 20)),
		validation.Field(&u.Password, validation.Required, validation.Length(4, 32)),
		validation.Field(&u.UserType, validation.Required, validation.In("Учитель", "Ученик", "Родитель")),
	)
}

func (u User) ValidateForLogin() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Login, validation.Required, validation.Length(4, 20)),
		validation.Field(&u.Password, validation.Required, validation.Length(4, 32)),
	)
}
