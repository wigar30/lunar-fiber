package entity

func (s *Membership) TableName() string {
	return "memberships"
}

type Membership struct {
	ID            string  `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID        string  `json:"userId,omitempty" gorm:"not null;column:userId"`
	User          *User   `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID"`
	RoleID        string  `json:"roleId,omitempty" gorm:"not null;column:roleId"`
	Role          *Role   `json:"role,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID"`
	TenantID      string  `json:"-" gorm:"not null;column:tenantId"`
	Tenant        *Tenant `json:"tenant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	StatusID      string  `json:"-" gorm:"not null;column:statusId"`
	Status        *Status `json:"status,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:StatusID"`
	DefaultColumn `gorm:"embedded"`
}
