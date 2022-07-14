package entitysumm

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&Location{})

	if err != nil {
		panic("Could not migrate sumArea")
	}

	err = db.AutoMigrate(&SensorSummarizeArea{})

	if err != nil {
		panic("Could not migrate SensorSummarizeArea")
	}

	err = db.AutoMigrate(&SumArea{})

	if err != nil {
		panic("Could not migrate summarize areas")
	}

	err = db.AutoMigrate(&SumAreaAssignment{})

	if err != nil {
		panic("Could not migrate SumAreaAssignment")
	}

	err = db.AutoMigrate(&Sensor{})

	if err != nil {
		panic("Could not migrate Sensor")
	}
}
