package entity

func (u *User) TableName() string {
	return "users"
}

type User struct {
	ID       string  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email    string  `json:"email,omitempty" gorm:"unique;not null"`
	Password string  `json:"-"`
	Name     string  `json:"name,omitempty"`
	RoleID   string  `json:"roleId,omitempty" gorm:"not null;column:roleId"`
	Role     *Role   `json:"role,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID"`
	StatusID string  `json:"statusId,omitempty" gorm:"not null;column:statusId"`
	Status   *Status `json:"status,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:StatusID"`
	Hash     string  `json:"hash,omitempty"`
	// Memberships   []Membership `json:"membership,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserRefer"`
	DefaultColumn `gorm:"embedded"`
}
