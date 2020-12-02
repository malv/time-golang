package model

// LeaveRequest ...
type LeaveRequest struct {
	BaseModel
	LeaveID     string `gorm:"column:leave_id" json:"leaveId"`
	Leave       Leave  `gorm:"foreignKey:LeaveID"`
	EmployeeID  string `gorm:"column:employee_id" json:"employeeId"`
	StartDate   Date   `gorm:"start_date;type:date" json:"startDate"`
	EndDate     Date   `gorm:"end_date;type:date" json:"endDate"`
	Description string `gorm:"description" json:"description"`
	Status      string `gorm:"status" json:"status"`
}

// TableName ...
func (LeaveRequest) TableName() string {
	return "tim_leave_requests"
}
