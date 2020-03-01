package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)
import "github.com/satori/go.uuid"

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open("postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		),
	)
}
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	xuuid := uuid.NewV4()
	return scope.SetColumn("Id", xuuid.String())
}
