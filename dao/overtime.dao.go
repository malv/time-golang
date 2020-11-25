package dao

import (
	"time-project/config"
	"time-project/model"
)

// OvertimeDao ...
type OvertimeDao struct{}

// GetOvertimes ...
func (OvertimeDao) GetOvertimes() (data []model.Overtime, e error) {
	defer config.CatchError(&e)

	var overtime []model.Overtime

	result := g.Find(&overtime)

	if result.Error != nil {
		return data, result.Error
	}

	return overtime, nil
}

// GetOvertimeByID ...
func (OvertimeDao) GetOvertimeByID(id string) (data model.Overtime, e error) {
	defer config.CatchError(&e)

	overtime := model.Overtime{BaseModel: model.BaseModel{ID: id}}

	result := g.First(&overtime)

	if result.Error != nil {
		return data, result.Error
	}

	return overtime, nil
}

// AddOvertime ...
func (OvertimeDao) AddOvertime(data *model.Overtime) (e error) {
	defer config.CatchError(&e)

	result := g.Create(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// EditOvertime ...
func (OvertimeDao) EditOvertime(data *model.Overtime) (e error) {
	defer config.CatchError(&e)

	result := g.Save(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteOvertime ...
func (OvertimeDao) DeleteOvertime(id string) (e error) {
	defer config.CatchError(&e)
	var overtime model.Overtime
	result := g.Delete(overtime, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
