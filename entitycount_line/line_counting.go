package entitycount_line

import (
	"measurement_validation_service/objects"
	"time"

	"gorm.io/gorm"
)

type LineCountingSensor struct {
	Id           uint                 `mapper:"_id"`
	Name         string               `mapper:"_Name"`
	CountingLine objects.GormGeometry ` mapper:"-"`
	LocationId   *uint                `mapper:"_locationId"`
	Active       bool                 `mapper:"_active"`
	ResetTime    time.Time
	ExternalId   string ` mapper:"_externalId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
