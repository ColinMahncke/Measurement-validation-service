package entitysumm

import "time"

type SumFlowValues struct {
	Timestamp time.Time `gorm:"primarykey" mapper:"_timestamp"`
	Unit      string    `gorm:"primarykey" mapper:"_unit"`
	SumAreaId uint      `gorm:"primarykey" mapper:"_sumAreaId"`
	SumArea   SumArea   `mapper:"_" gorm:"foreignKey:SensorId;"`
	OutValue  int       `mapper:"_outValue"`
	InValue   int       `mapper:"_inValue"`
}
