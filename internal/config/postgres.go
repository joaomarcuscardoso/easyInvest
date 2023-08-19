package config

import (
	"fmt"

	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	fmt.Println("Initalize Postgres")

	// logger := GetLogger("postgres")
	// dbPath := "host=localhost user=postgres password=postgres dbname=easy_invest port=5432 sslmode=disable TimeZone=Brazil/East"

	return nil, nil
}
