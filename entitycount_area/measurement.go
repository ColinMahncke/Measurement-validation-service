package entitycount_area

import (
	"time"
)

type AreaCountingSensorMeasurement struct {
	AreaCountingSensorId  uint                           `gorm:"primaryKey;autoIncrement:false"`
	AreaCountingSensor    AreaCountingSensor             `gorm:"foreignKey:AreaCountingSensorId;"`
	Timestamp             time.Time                      `gorm:"primaryKey" mapper:"_measurementTimestamp"`
	Unit                  string                         `gorm:"primaryKey" mapper:"_unit"`
	Value                 int                            `mapper:"_value"`
	ExternalId            string                         `gorm:"-" mapper:"_externalId"`
	OverriddenMeasurement *AreaCountingSensorMeasurement `gorm:"-"`
}
