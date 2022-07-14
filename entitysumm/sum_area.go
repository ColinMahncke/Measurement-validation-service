package entitysumm

import (
	"measurement_validation_service/objects"
	"time"

	"gorm.io/gorm"
)

type SumArea struct {
	Id            uint                 `gorm:"primarykey" mapper:"_id"`
	Name          string               `mapper:"_name"`
	AreaType      string               `mapper:"_area_type"`
	LocationId    *uint                `mapper:"_locationId" gorm:"index"`
	Location      Location             `mapper:"-" gorm:"foreignKey:LocationId;constraint:OnDelete:SET NULL"`
	MeasuringZone objects.GormGeometry `mapper:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
