package main

import (
	"database/sql"
	"fmt"
	"measurement_validation_service/entitycount_area"
	"measurement_validation_service/entitycount_line"
	"measurement_validation_service/entitysumm"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", Getenv("host", "host"), Getenv("user", "user"), Getenv("password", "password"), Getenv("dbname", "dbname"), Getenv("port", "port"), Getenv("sslmode", "sslmode"), Getenv("Timezone", "Timezone"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	dsnLine := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", Getenv("host", "host"), Getenv("user", "user"), Getenv("password", "password"), Getenv("dbnameline", "dbnameline"), Getenv("port", "port"), Getenv("sslmode", "sslmode"), Getenv("Timezone", "Timezone"))
	dbLine, err := gorm.Open(postgres.Open(dsnLine), &gorm.Config{})
	if err != nil {
		return
	}

	dsnCount := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", Getenv("host", "host"), Getenv("user", "user"), Getenv("password", "password"), Getenv("dbnamearea", "dbnamearea"), Getenv("port", "port"), Getenv("sslmode", "sslmode"), Getenv("Timezone", "Timezone"))
	dbCount, err := gorm.Open(postgres.Open(dsnCount), &gorm.Config{})
	if err != nil {
		return
	}
	entitycount_area.AutoMigrate(dbCount)

	entitycount_line.AutoMigrate(dbLine)

	entitysumm.AutoMigrate(db)

	offset := 0
	for {
		var summvalues entitysumm.SumValues

		if result := db.Limit(1).Offset(offset).Find(&summvalues); result.RowsAffected < 1 {
			break
		}
		sensorsInArea := addSensorsInArea(summvalues.SumAreaId, db)
		totalSensorMeasurement := 0

		for _, sensorsInArea := range sensorsInArea {
			//fmt.Println(fmt.Sprintf("Id: %d, SensorId: %d, SensorType: %s", sensorsInArea.Id, sensorsInArea.Sensor.ExternSensorId, sensorsInArea.Sensor.SensorType))
			if sensorsInArea.Sensor.SensorType == "area_counting" {
				var measurement entitycount_area.AreaCountingSensorMeasurement
				dbCount.Order("timestamp desc").Where("timestamp <= ? and area_counting_sensor_id=? and unit=?", summvalues.Timestamp, sensorsInArea.Sensor.ExternSensorId, summvalues.Unit).First(&measurement)

				//fmt.Println(measurement.Value)
				totalSensorMeasurement = totalSensorMeasurement + measurement.Value
			} else if sensorsInArea.Sensor.SensorType == "line_counting" {
				var totalValue *int

				var line_sensor entitycount_line.LineCountingSensor
				dbLine.Find(&line_sensor, "id =?", sensorsInArea.Sensor.ExternSensorId)
				last_reset := time.Date(summvalues.Timestamp.Year(), summvalues.Timestamp.Month(), summvalues.Timestamp.Day(), line_sensor.ResetTime.Hour(), line_sensor.ResetTime.Minute(), line_sensor.ResetTime.Second(), line_sensor.ResetTime.Nanosecond(), time.UTC)
				if last_reset.After(summvalues.Timestamp) {
					last_reset = last_reset.AddDate(0, 0, -1)
				}

				dbLine.Debug().Raw(`select sum(lcsm.in_value) - sum(lcsm.out_value) totalValue from public.line_counting_sensor_measurements lcsm where lcsm.measurement_timestamp  > @reset_timestamp and lcsm.measurement_timestamp <= @measurementTimestamp and lcsm.unit = @unit and lcsm.line_counting_sensor_id = @sensorId`, sql.Named("sensorId", sensorsInArea.Sensor.ExternSensorId), sql.Named("measurementTimestamp", summvalues.Timestamp), sql.Named("unit", summvalues.Unit), sql.Named("reset_timestamp", last_reset)).Scan(&totalValue)
				fmt.Println(totalValue)

				if totalValue != nil {
					totalSensorMeasurement += *totalValue
				}
			}

		}

		fmt.Println(fmt.Sprintf("AreaId: %d, Unit: %s, Timestamp: %s, Value: %d, ExpectedValue: %d, Correct: %t",
			summvalues.SumAreaId, summvalues.Unit, summvalues.Timestamp.Format("2006-01-02 15:04:05"), summvalues.Value, totalSensorMeasurement, totalSensorMeasurement == summvalues.Value))

		offset++

	}
}
func addSensorsInArea(sumAreaId uint, db *gorm.DB) []entitysumm.SensorSummarizeArea {
	var sensorsInArea []entitysumm.SensorSummarizeArea

	db.Preload("Sensor").Find(&sensorsInArea, "sum_area_id = ?", sumAreaId)

	var areaInArea []entitysumm.SumAreaAssignment

	db.Find(&areaInArea, "parent_id = ?", sumAreaId)

	for _, areaAssignment := range areaInArea {
		addSensorsInArea(areaAssignment.SumAreaId, db)
		sensorsInArea = append(sensorsInArea, addSensorsInArea(areaAssignment.SumAreaId, db)...)
	}

	return sensorsInArea

}
