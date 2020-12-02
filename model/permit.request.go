package model

// PermitRequest ...
type PermitRequest struct {
	BaseModel
	PermitID    string `gorm:"column:permit_id" json:"permitId"`
	Permit      Permit `gorm:"foreignKey:PermitID"`
	EmployeeID  string `gorm:"column:employee_id" json:"employeeId"`
	Date        Date   `gorm:"column:date; type:date" json:"date"`
	Description string `gorm:"column:description" json:"description"`
	Status      string `gorm:"column:status" json:"status"`
}

// TableName ...
func (PermitRequest) TableName() string {
	return "tim_permit_requests"
}
