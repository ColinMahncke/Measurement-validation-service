package entitycount_area

import (
	"measurement_validation_service/objects"
	"time"

	"gorm.io/gorm"
)

type AreaCountingSensor struct {
	ID            uint `gorm:"primarykey" mapper:"_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt       `gorm:"index"`
	Name          string               `mapper:"_name"`
	MeasuringZone objects.GormGeometry `mapper:"-"`
	Position      objects.GormGeometry `mapper:"-"`
	Active        bool                 `mapper:"_active"`
	ExternalId    string               `mapper:"_externalId"`
	LocationId    *uint                `mapper:"_locationId" gorm:"index"`
	Location      Location             `mapper:"-" gorm:"foreignKey:LocationId;constraint:OnDelete:SET NULL"`
}
