package db_test

import (
	"fmt"
	"pishikot/pkg/models"
	"pishikot/pkg/store/db"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBConn() *gorm.DB {
	config, err := models.LoadConfig()
	if err != nil {
		panic(err)
	}
	_db, err := gorm.Open(postgres.Open(fmt.Sprintf("database=%s user=%s password=%s host=%s sslmode=%s",
		config.Database.Dbname, config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Sslmode)),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	return _db
}
func TestUserRepo_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		arg     *models.User
		wantErr bool
	}{
		{
			"Add user #1",
			&models.User{
				Login:    "testUser",
				Password: "testUser",
			},
			false,
		},
		{
			"Add user #2",
			&models.User{
				Login:    "sdffdsdf",
				Password: "fdgdwfgfds",
			},
			false,
		},
	}
	r := db.NewUserRepo(getDBConn())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.CreateUser(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if err := r.DeleteUser(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
