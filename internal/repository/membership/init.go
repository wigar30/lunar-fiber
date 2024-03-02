package membership

import (
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/model"
)

type MembershipRepository struct {
	db *driver.Database
}

func NewMembershipRepository(db *driver.Database) model.MembershipRepositoryInterface {
	return &MembershipRepository{
		db: db,
	}
}