package entitysumm

import (
	"time"

	"gorm.io/gorm"
)

type SensorSummarizeArea struct {
	Id                   uint      `gorm:"primarykey"`
	SensorId             uint      `mapper:"_"`
	Sensor               Sensor    `mapper:"_" gorm:"foreignKey:SensorId;"`
	SumAreaId            uint      `mapper:"_sumAreaId"`
	SumArea              SumArea   `mapper:"-" gorm:"foreignKey:SumAreaId;"`
	ReverseFlowDirection bool      `mapper:"_reverseFlowDirection"`
	CreatedAt            time.Time `mapper:"-"`
	DeletedAt            gorm.DeletedAt
}
