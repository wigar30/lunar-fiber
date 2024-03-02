package model

import (
	"lunar-commerce-fiber/internal/entity"

	"gorm.io/gorm"
)

type CreateMembership struct {
	UserID   string `json:"userId,omitempty" gorm:"not null;column:userId" validate:"required"`
	RoleID   string `json:"roleId,omitempty" gorm:"not null;column:roleId" validate:"required"`
	TenantID string `json:"-" gorm:"not null;column:tenantId" validate:"required"`
	StatusID string `json:"-" gorm:"not null;column:statusId" validate:"required"`
}

type MembershipsResponse struct {
	ID       string         `json:"id"`
	UserID   string         `json:"userId,omitempty"`
	User     *entity.User   `json:"user,omitempty"`
	RoleID   string         `json:"roleId,omitempty"`
	Role     *entity.Role   `json:"role,omitempty"`
	TenantID string         `json:"-"`
	Tenant   *entity.Tenant `json:"tenant,omitempty"`
	StatusID string         `json:"-"`
	Status   *entity.Status `json:"status,omitempty"`
}

type MembershipRepositoryInterface interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB) error
	RollbackTransaction(*gorm.DB) error

	CreateMembership(*gorm.DB, CreateMembership) (string, error)
}

type MembershipUseCaseInterface interface {
}
