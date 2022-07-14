package entitysumm

import "time"

type SumValues struct {
	Timestamp time.Time `gorm:"primarykey" mapper:"_timestamp"`
	Unit      string    `gorm:"primarykey" mapper:"_unit"`
	SumAreaId uint      `gorm:"primarykey" mapper:"_sumAreaId"`
	SumArea   SumArea   `mapper:"_" gorm:"foreignKey:SensorId;"`
	Value     int       `mapper:"_value"`
}
