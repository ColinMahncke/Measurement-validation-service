package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", Getenv("host", "host"), Getenv("user", "user"), Getenv("password", "password"), Getenv("dbname", "dbname"), Getenv("port", "port"), Getenv("sslmode", "sslmode"), Getenv("Timezone", "Timezone"))
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}
