package entity

func (u *SummaryStat) TableName() string {
	return "summary_stats"
}

type SummaryStat struct {
	ID                 string  `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID           string  `json:"tenantId,omitempty" gorm:"not null;column:tenantId"`
	Tenant             *Tenant `json:"tenant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TenantID"`
	UnprocessedOrder   int     `json:"unprocessed_order" gorm:"not null"`
	CompletedOrder     int     `json:"completed_order" gorm:"not null"`
	OrderBeingSent     int     `json:"order_being_sent" gorm:"not null"`
	UnfinishedComplain int     `json:"unfinished_complain" gorm:"not null"`
	TotalComplain      int     `json:"total_complain" gorm:"not null"`
	DefaultColumn      `gorm:"embedded"`
}
