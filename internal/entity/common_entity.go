package entity

import (
	"time"

	"gorm.io/gorm"
)

type DefaultColumn struct {
	CreatedAt *time.Time      `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt *time.Time      `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deletedAt"`
}
