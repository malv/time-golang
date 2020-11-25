package dao

import (
	"time-project/config"
	"time-project/model"
)

// PermitDao ...
type PermitDao struct{}

// GetPermits ...
func (PermitDao) GetPermits() (data []model.Permit, e error) {
	defer config.CatchError(&e)

	var permit []model.Permit

	result := g.Find(&permit)

	if result.Error != nil {
		return data, result.Error
	}

	return permit, nil
}

// GetPermitByID ...
func (PermitDao) GetPermitByID(id string) (data model.Permit, e error) {
	defer config.CatchError(&e)

	permit := model.Permit{BaseModel: model.BaseModel{ID: id}}

	result := g.First(&permit)

	if result.Error != nil {
		return data, result.Error
	}

	return permit, nil
}

// AddPermit ...
func (PermitDao) AddPermit(data *model.Permit) (e error) {
	defer config.CatchError(&e)

	result := g.Create(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// EditPermit ...
func (PermitDao) EditPermit(data *model.Permit) (e error) {
	defer config.CatchError(&e)

	result := g.Save(data)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeletePermit ...
func (PermitDao) DeletePermit(id string) (e error) {
	defer config.CatchError(&e)

	var permit model.Permit

	result := g.Delete(&permit, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
