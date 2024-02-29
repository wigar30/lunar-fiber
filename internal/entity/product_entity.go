package entity

import "gorm.io/gorm"

func (s *Product) TableName() string {
	return "products"
}

type Product struct {
	ID            string  `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID      string  `json:"tenantId,omitempty" gorm:"not null;column:tenantId"`
	Name          string  `json:"name" gorm:"not null"`
	TotalStock    *int    `json:"totalStock,omitempty" gorm:"not null;column:totalStock;default:0"`
	TotalSold     *int    `json:"totalSold,omitempty" gorm:"not null;column:totalSold;default:0"`
	Price         *int    `json:"price,omitempty" gorm:"not null;column:price;default:0"`
	Description   string  `json:"description,omitempty" gorm:"column:description;"`
	Specification string  `json:"specification,omitempty" gorm:"column:specification;"`
	Tenant        *Tenant `json:"tenant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	DefaultColumn `gorm:"embedded"`
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	var tenant Tenant
	err = tx.Select("id", "total_product").First(&tenant, p.TenantID).Error
	if err != nil {
		return err
	}
	tx.Table("Tenant").Where("id", p.TenantID).Update("total_product", tenant.TotalProduct + 1)

	return
}
