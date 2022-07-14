package entitycount_area

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&Location{}, &AreaCountingSensor{}, &AreaCountingSensorMeasurement{})

	if err != nil {
		panic(err)
	}
}
