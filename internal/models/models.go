package models

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBModels struct {
	Order OrderModel
	User UserModel
	DB *gorm.DB
}

func InitDB(dataSourceName string) (*DBModels, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to Migrate: %v", err)
	}
	err = db.AutoMigrate(&Order{}, &OrderItem{}, &User{})
	if err != nil {
		return nil, fmt.Errorf("Failed to Migrate: %v", err)
	}

	dbModel := &DBModels{
		DB: db,
		Order: OrderModel{DB: db},
		User: UserModel{DB: db},
	}

	return dbModel, nil

}