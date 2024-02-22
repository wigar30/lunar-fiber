package entity

func (u *Tenant) TableName() string {
	return "tenants"
}

type Tenant struct {
	ID            string       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string       `json:"name" gorm:"not null"`
	TotalProduct  *int         `json:"totalProduct,omitempty" gorm:"not null;column:totalProduct;default:0"`
	LevelID       string       `json:"levelId,omitempty" gorm:"not null;column:levelId"`
	LevelTenant   LevelTenant  `json:"levelTenant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:LevelID"`
	Memberships   []Membership `json:"memberships,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	Products      []*Product   `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID;embeddedPrefix:postal_address_"`
	DefaultColumn `gorm:"embedded"`
}
