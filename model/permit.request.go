package model

// PermitRequest ...
type PermitRequest struct {
	BaseModel
	PermitID    string    `gorm:"column:permit_id" json:"permitId"`
	Permit      Permit    `gorm:"foreignKey:PermitID"`
	Date        Timestamp `gorm:"column:date" json:"date"`
	Description string    `gorm:"column:description" json:"description"`
	Status      string    `gorm:"column:status" json:"status"`
}

// TableName ...
func (PermitRequest) TableName() string {
	return "tim_permit_requests"
}
