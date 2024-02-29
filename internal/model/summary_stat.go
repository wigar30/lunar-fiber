package model

import "lunar-commerce-fiber/internal/entity"

type SummaryStatResponse struct {
	ID                 string         `json:"id"`
	TenantID           string         `json:"-"`
	Tenant             *entity.Tenant `json:"tenant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	UnprocessedOrder   int64          `json:"unprocessedOrder"`
	CompletedOrder     int64          `json:"completedOrder"`
	OrderBeingSent     int64          `json:"orderBeingSent"`
	UnfinishedComplain int64          `json:"unfinishedComplain"`
	TotalComplain      int64          `json:"totalComplain"`
}
