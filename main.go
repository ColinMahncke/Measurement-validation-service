package main

import (
	"fmt"
	"measurement_validation_service/entitysumm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", Getenv("host", "host"), Getenv("user", "user"), Getenv("password", "password"), Getenv("dbname", "dbname"), Getenv("port", "port"), Getenv("sslmode", "sslmode"), Getenv("Timezone", "Timezone"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	//entitycount_area.AutoMigrate(db)

	//entitycount_line.AutoMigrate(db)

	entitysumm.AutoMigrate(db)

	var summvalues entitysumm.SumValues
	db.Limit(1).Offset(0).Preload("SumArea").Find(&summvalues)

}
