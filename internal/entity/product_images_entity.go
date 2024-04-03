package entity

func (ProductImage) TableName() string {
	return "product_images"
}

type ProductImage struct {
	ID            string   `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID      string   `json:"tenantId,omitempty" gorm:"not null;column:tenantId"`
	ProductID     string   `json:"productId,omitempty" gorm:"not null;column:productId"`
	StatusID      string   `json:"statusId,omitempty" gorm:"not null;column:statusId"`
	Title         string   `json:"title,omitempty" gorm:"column:title;"`
	Image         string   `json:"image" gorm:"column:image;"`
	Tenant        *Tenant  `json:"tenant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	Product       *Product `json:"product,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:ProductID"`
	Status        *Status  `json:"status,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:StatusID"`
	DefaultColumn `gorm:"embedded"`
}
