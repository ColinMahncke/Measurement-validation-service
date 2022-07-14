package entitysumm

import (
	"gorm.io/gorm"

	"time"
)

type SumAreaAssignment struct {
	Id            uint      `gorm:"primarykey"`
	ParentId      uint      `mapper:"_parentId"`
	SumAreaId     uint      `mapper:"_sumAreaId"`
	SumArea       SumArea   `mapper:"-" gorm:"foreignKey:SumAreaId;"`
	SumAreaParent SumArea   `mapper:"-" gorm:"foreignKey:ParentId;"`
	CreatedAt     time.Time `mapper:"-"`
	DeletedAt     gorm.DeletedAt
}
