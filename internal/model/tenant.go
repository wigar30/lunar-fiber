package model

import (
	"lunar-commerce-fiber/internal/entity"
	"lunar-commerce-fiber/pkg/utils"

	"gorm.io/gorm"
)

type CreateTenant struct {
	Name string `json:"name" gorm:"not null" validate:"required"`
}

type TenantResponse struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	TotalProduct int64               `json:"total_product"`
	LevelID      string              `json:"levelId"`
	Level        *entity.LevelTenant `json:"level,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:LevelID"`
	SummaryStat  *entity.SummaryStat `json:"summaryStat" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	Memberships  []entity.Membership `json:"memberships" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
}

type TenantRepositoryInterface interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB) error
	RollbackTransaction(*gorm.DB) error

	GetAllByAuth(int64, utils.Pagination) (*utils.Pagination, error)
	GetCountAuthTenant(int64) (int64, error)
	GetByID(string, string) (*entity.Tenant, error)
	CreateTenant(*gorm.DB, CreateTenant) (string, error)
}

type TenantUseCaseInterface interface {
	GetAllByAuth(int64, PaginationRequest) (*utils.Pagination, error)
	GetByID(string, string) (*TenantResponse, error)
	CreateTenant(int64, CreateTenant) (string, error)
}
