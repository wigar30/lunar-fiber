package entity

func (s *Product) TableName() string {
	return "products"
}

type Product struct {
	ID            string  `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantRefer   string  `json:"tenantId,omitempty" gorm:"not null;column:tenantId"`
	Name          string  `json:"name" gorm:"not null"`
	TotalStock    *int    `json:"totalStock,omitempty" gorm:"not null;column:totalStock;default:0"`
	TotalSold     *int    `json:"totalSold,omitempty" gorm:"not null;column:totalSold;default:0"`
	Price         *int    `json:"price,omitempty" gorm:"not null;column:price;default:0"`
	Description   string  `json:"description,omitempty" gorm:"column:description;"`
	Specification string  `json:"specification,omitempty" gorm:"column:specification;"`
	Tenant        *Tenant `json:"tenant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantRefer"`
	DefaultColumn `gorm:"embedded"`
}
