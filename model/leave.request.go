package model

// LeaveRequest ...
type LeaveRequest struct {
	BaseModel
	LeaveID     string    `gorm:"column:leave_id"`
	Leave       Leave     `gorm:"foreignKey:LeaveID"`
	StartDate   Timestamp `gorm:"start_date" json:"startDate"`
	EndDate     Timestamp `gorm:"end_date" json:"endDate"`
	Description string    `gorm:"description" json:"description"`
	Status      string    `gorm:"status" json:"status"`
}

// TableName ...
func (LeaveRequest) TableName() string {
	return "tim_leave_requests"
}
