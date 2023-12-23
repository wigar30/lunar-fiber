package entity

func (s *Membership) TableName() string {
	return "memberships"
}

type Membership struct {
	ID            string  `json:"id" gorm:"primaryKey;autoIncrement"`
	UserRefer     string  `json:"userId,omitempty" gorm:"not null;column:userId"`
	User          *User   `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserRefer"`
	RoleID        string  `json:"roleId,omitempty" gorm:"not null;column:roleId"`
	Role          *Role   `json:"role,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID"`
	TenantRefer   string  `json:"-" gorm:"not null;type:uuid;column:tenantId"`
	Tenant        *Tenant `json:"tenant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantRefer"`
	StatusID      string  `json:"-" gorm:"not null;column:statusId"`
	Status        *Status `json:"status,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:StatusID"`
	DefaultColumn `gorm:"embedded"`
}
