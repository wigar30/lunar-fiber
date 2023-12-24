package model

import "lunar-commerce-fiber/internal/entity"

type UserResponse struct {
	ID       string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email    string         `json:"email,omitempty" gorm:"unique;not null"`
	Name     string         `json:"name,omitempty"`
	RoleID   string         `json:"roleId,omitempty" gorm:"not null;column:roleId"`
	Role     *entity.Role   `json:"role,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID"`
	StatusID string         `json:"statusId,omitempty" gorm:"not null;column:statusId"`
	Status   *entity.Status `json:"status,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:StatusID"`
}

type UserRepositoryInterface interface {
	GetUserByEmail(string) (*entity.User, error)
}
