package dao

import (
	"time-project/config"
	"time-project/model"
)

// LeaveDao ...
type LeaveDao struct{}

// GetLeaves ...
func (LeaveDao) GetLeaves() (data []model.Leave, e error) {
	defer config.CatchError(&e)
	var leaves []model.Leave
	result := g.Find(&leaves)
	if result.Error == nil {
		return leaves, nil
	}
	return data, result.Error
}

// GetLeaveByID ...
func (LeaveDao) GetLeaveByID(id string) (data model.Leave, e error) {
	defer config.CatchError(&e)
	leave := model.Leave{BaseModel: model.BaseModel{ID: id}}
	result := g.First(&leave)
	if result.Error == nil {
		return leave, nil
	}
	return data, result.Error
}

// AddLeave ...
func (LeaveDao) AddLeave(data *model.Leave) (e error) {
	defer config.CatchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

// EditLeave ...
func (LeaveDao) EditLeave(data *model.Leave) (e error) {
	defer config.CatchError(&e)
	result := g.Save(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

// DeleteLeave ...
func (LeaveDao) DeleteLeave(id string) (e error) {
	defer config.CatchError(&e)
	var leave model.Leave
	result := g.Delete(&leave, id)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
