package model

// OvertimeRequest ...
type OvertimeRequest struct {
	BaseModel
	OvertimeID  string    `gorm:"column:overtime_id" json:"overtimeId"`
	Overtime    Overtime  `gorm:"foreignKey:OvertimeID"`
	EmployeeID  string    `gorm:"column:employee_id" json:"employeeId"`
	StartTime   Timestamp `gorm:"column:start_time" json:"startTime"`
	EndTime     Timestamp `gorm:"column:end_time" json:"endTime"`
	Description string    `gorm:"column:description" json:"description"`
	Status      string    `gorm:"column:status" json:"status"`
}

// TableName ...
func (OvertimeRequest) TableName() string {
	return "tim_overtime_requests"
}
