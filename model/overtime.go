package model

// Overtime ...
type Overtime struct {
	BaseModel
	Name        string `gorm:"column:name" json:"name"`
	CompanyID   string `gorm:"column:company_id" json:"companyId"`
	Code        string `gorm:"column:code" json:"code"`
	Description string `gorm:"column:description" json:"description"`
}

// TableName ...
func (Overtime) TableName() string {
	return "tim_overtimes"
}
