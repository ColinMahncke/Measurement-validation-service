package entitycount_line

import (
	"time"
)

// LineCountingSensorMeasurement todo fix foreign Key
type LineCountingSensorMeasurement struct {
	LineCountingSensorId uint               `gorm:"primaryKey;autoIncrement:false;index"`
	Sensor               LineCountingSensor `gorm:"foreignKey:LineCountingSensorId"`
	MeasurementTimestamp time.Time          `gorm:"primaryKey" mapper:"_measurementTimestamp"`
	Unit                 string             `gorm:"primaryKey" mapper:"_unit"`
	OutValue             int                `mapper:"-"`
	InValue              int                `mapper:"-"`
	OutDiv               int                `gorm:"-" mapper:"-"`
	InDiv                int                `gorm:"-" mapper:"-"`
}
