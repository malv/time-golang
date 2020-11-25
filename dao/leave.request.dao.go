package dao

import (
	"time-project/config"
	"time-project/model"
)

// LeaveRequestDao ...
type LeaveRequestDao struct{}

// GetLeaveRequests ...
func (LeaveRequestDao) GetLeaveRequests() (data []model.LeaveRequest, e error) {
	defer config.CatchError(&e)

	var leaveRequest []model.LeaveRequest

	result := g.Find(&leaveRequest)

	if result.Error != nil {
		return data, result.Error
	}
	return leaveRequest, nil
}

// GetLeaveRequestByID ...
func (LeaveRequestDao) GetLeaveRequestByID(id string) (data model.LeaveRequest, e error) {
	defer config.CatchError(&e)

	leaveRequest := model.LeaveRequest{BaseModel: model.BaseModel{ID: id}}

	result := g.First(&leaveRequest)

	if result.Error != nil {
		return data, result.Error
	}
	return leaveRequest, nil
}

// AddLeaveRequest ...
func (LeaveRequestDao) AddLeaveRequest(data *model.LeaveRequest) (e error) {
	defer config.CatchError(&e)

	result := g.Create(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// EditLeaveRequest ...
func (LeaveRequestDao) EditLeaveRequest(data *model.LeaveRequest) (e error) {
	defer config.CatchError(&e)

	result := g.Save(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteLeaveRequest ...
func (LeaveRequestDao) DeleteLeaveRequest(id string) (e error) {
	defer config.CatchError(&e)

	var leaveRequest model.LeaveRequest

	result := g.Delete(&leaveRequest, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
