package entity

func (s *Status) TableName() string {
	return "statuses"
}

type Status struct {
	ID            string `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string `json:"name" gorm:"unique;not null"`
	DefaultColumn `gorm:"embedded"`
}