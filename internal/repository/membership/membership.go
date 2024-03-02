package membership

import (
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
