package entity

type Tenant struct {
	ID            string       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string       `json:"name" gorm:"not null"`
	TotalProduct  int          `json:"totalProduct,omitempty" gorm:"not null;column:totalProduct;default:0"`
	LevelID       string       `json:"levelId,omitempty" gorm:"not null;column:levelId"`
	LevelTenant   LevelTenant  `json:"levelTenant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:LevelID"`
	Memberships   []Membership `json:"memberships,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	Products      []*Product   `json:"products,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	SummaryStat   *SummaryStat `json:"summaryStat" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	DefaultColumn `gorm:"embedded"`
}
