package entity

import (
	"time"

	"gorm.io/gorm"
)

type DefaultColumn struct {
	CreatedAt *time.Time      `json:"-" gorm:"column:createdAt"`
	UpdatedAt *time.Time      `json:"-" gorm:"column:updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deletedAt"`
}
