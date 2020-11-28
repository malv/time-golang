package dao

import (
	"time-project/config"
	"time-project/model"
)

// OvertimeRequestDao ...
type OvertimeRequestDao struct{}

// GetOvertimeRequests ...
func (OvertimeRequestDao) GetOvertimeRequests() (data []model.OvertimeRequest, e error) {
	defer config.CatchError(&e)

	var overtimeRequest []model.OvertimeRequest

	result := g.Find(&overtimeRequest)

	if result.Error != nil {
		return data, result.Error
	}
	return overtimeRequest, nil
}

// GetOvertimeRequestByID ...
func (OvertimeRequestDao) GetOvertimeRequestByID(id string) (data model.OvertimeRequest, e error) {
	defer config.CatchError(&e)

	overtimeRequest := model.OvertimeRequest{BaseModel: model.BaseModel{ID: id}}

	result := g.First(&overtimeRequest)

	if result.Error != nil {
		return data, result.Error
	}
	return overtimeRequest, nil
}

// AddOvertimeRequest ...
func (OvertimeRequestDao) AddOvertimeRequest(data *model.OvertimeRequest) (e error) {
	defer config.CatchError(&e)

	result := g.Create(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// EditOvertimeRequest ...
func (OvertimeRequestDao) EditOvertimeRequest(data *model.OvertimeRequest) (e error) {
	defer config.CatchError(&e)

	result := g.Save(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteOvertimeRequest ...
func (OvertimeRequestDao) DeleteOvertimeRequest(id string) (e error) {
	defer config.CatchError(&e)

	var overtimeRequest model.OvertimeRequest

	result := g.Delete(&overtimeRequest, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
