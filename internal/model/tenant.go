package model

import (
	"lunar-commerce-fiber/internal/entity"
	"lunar-commerce-fiber/pkg/utils"
)

type TenantResponse struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	TotalProduct int64               `json:"total_product"`
	LevelID      string              `json:"levelId"`
	Level        *entity.LevelTenant `json:"level,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:LevelID"`
	Memberships  []entity.Membership `json:"memberships" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
}

type TenantRepositoryInterface interface {
	GetAllByAuth(int64, utils.Pagination) (*utils.Pagination, error)
	GetByID(string, string) (*entity.Tenant, error)
}

type TenantUseCaseInterface interface {
	GetAllByAuth(int64, PaginationRequest) (*utils.Pagination, error)
	GetByID(string, string) (*TenantResponse, error)
}
