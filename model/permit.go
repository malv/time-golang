package model

// Permit ...
type Permit struct {
	BaseModel
	Name        string `gorm:"column:name" json:"name"`
	CompanyID   string `gorm:"column:company_id" json:"companyId"`
	Code        string `gorm:"column:code" json:"code"`
	Description string `gorm:"column:description" json:"description"`
}

// TableName ...
func (Permit) TableName() string {
	return "tim_permits"
}
