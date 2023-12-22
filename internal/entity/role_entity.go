package entity

func (r *Role) TableName() string {
	return "roles"
}

type Role struct {
	ID            string  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string  `json:"name" gorm:"unique;not null"`
	Users         *[]User `json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DefaultColumn `gorm:"embedded"`
}
