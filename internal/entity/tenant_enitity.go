package entity

func (u *Tenant) TableName() string {
	return "tenants"
}

type Tenant struct {
	ID            string       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string       `json:"name" gorm:"not null"`
	TotalProduct  *int         `json:"totalProduct,omitempty" gorm:"not null;column:totalProduct;default:0"`
	Memberships   []Membership `json:"memberships,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantRefer"`
	Products      []*Product   `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantRefer;embeddedPrefix:postal_address_"`
	DefaultColumn `gorm:"embedded"`
}
