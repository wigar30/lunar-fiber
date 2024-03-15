package membership

import (
	"errors"
	"lunar-commerce-fiber/internal/entity"
	"lunar-commerce-fiber/internal/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (mr *MembershipRepository) BeginTransaction() *gorm.DB {
	return mr.db.Begin()
}

func (mr *MembershipRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (mr *MembershipRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (tr *MembershipRepository) CheckIsAuthTenant(userID string, tenantID string) (bool, error) {
	var membership *entity.Membership
	
	err := tr.db.Model(&membership).Where(&entity.Membership{ UserID: userID, TenantID: tenantID }).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, &model.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		}
	}
	if err != nil {
		return false, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return true, nil
}

func (mr *MembershipRepository) CreateMembership(tx *gorm.DB, membership model.CreateMembership) (string, error) {
	membershipResult := &entity.Membership{
		UserID:   membership.UserID,
		TenantID: membership.TenantID,
		StatusID: membership.StatusID,
		RoleID:   membership.RoleID,
	}

	err := tx.Create(&membershipResult).Error
	if err != nil {
		return "", &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return membershipResult.ID, nil
}
