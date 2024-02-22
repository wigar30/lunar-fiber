package entity

func (s *LevelTenant) TableName() string {
	return "level_tenants"
}

type LevelTenant struct {
	ID            string `json:"id" gorm:"primaryKey;autoIncrement"`
	Level         string `json:"level" gorm:"not null"`
	DefaultColumn `gorm:"embedded"`
}
