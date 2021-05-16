package store

import (
	"fmt"
	"os"
	"pishikot/pkg/models"
	"pishikot/pkg/store/db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	Config   *models.Config
	UserRepo *db.UserRepo
	TaskRepo *db.TaskRepo
}

func New() *Store {
	config, err := models.LoadConfig()
	if err != nil {
		panic(err)
	}
	config.Database.Password = os.Getenv("DBPASSWORD")
	_db, err := gorm.Open(postgres.Open(fmt.Sprintf("database=%s user=%s password=%s host=%s sslmode=%s",
		config.Database.Dbname, config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Sslmode)),
		&gorm.Config{})
	_db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		panic(err)
	}
	return &Store{
		Config:   config,
		UserRepo: db.NewUserRepo(_db),
	}
}
