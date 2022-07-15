package entitysumm

import (
	"gorm.io/gorm"
)

type Sensor struct {
	Id             uint           `gorm:"primarykey" mapper:"_id"`
	SensorType     string         `mapper:"_sensorType"`
	ExternSensorId uint           `mapper:"_sensorId"`
	SensorName     string         `mapper:"_sensorName"`
	ExternalId     string         `mapper:"_externalId"`
	LocationId     uint           `mapper:"_locationId" gorm:"constraint:OnDelete:SET NULL"`
	DeletedAt      gorm.DeletedAt `mapper:"_deletedAt"`
}
