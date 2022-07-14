package entitycount_line

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {

	err := db.AutoMigrate(&LineCountingSensor{}, &Location{}, &LineCountingSensorMeasurement{})

	if err != nil {
		panic(err)
	}

}
