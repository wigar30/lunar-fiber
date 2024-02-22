package model

import "lunar-commerce-fiber/internal/entity"

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
