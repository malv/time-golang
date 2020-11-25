package model

// Leave ...
type Leave struct {
	BaseModel
	Name        string `gorm:"column:name" json:"name"`
	CompanyID   string `gorm:"column:company_id" json:"companyId"`
	Code        string `gorm:"column:code" json:"code"`
	Description string `gorm:"column:description" json:"description"`
}

// TableName ...
func (Leave) TableName() string {
	return "tim_leaves"
}
