package service

import (
	"time-project/config"
	"time-project/dao"
	"time-project/model"
)

// LeaveService ...
type LeaveService struct{}

var leaveDao dao.LeaveDao = dao.LeaveDao{}

// GetLeaves ...
func (LeaveService) GetLeaves() (data []model.Leave, e error) {
	defer config.CatchError(&e)
	result, err := leaveDao.GetLeaves()
	if err == nil {
		return result, nil
	}
	return data, err
}

// GetLeaveByID ...
func (LeaveService) GetLeaveByID(id string) (data model.Leave, e error) {
	defer config.CatchError(&e)
	result, err := leaveDao.GetLeaveByID(id)
	if err == nil {
		return result, nil
	}
	return data, err
}

// AddLeave ...
func (LeaveService) AddLeave(data *model.Leave) (e error) {
	defer config.CatchError(&e)
	data.CreatedBy = "Admin"
	admin := "Admin"
	data.UpdatedBy = &admin
	return leaveDao.AddLeave(data)
}

// EditLeave ...
func (LeaveService) EditLeave(data *model.Leave) (e error) {
	defer config.CatchError(&e)
	return leaveDao.EditLeave(data)
}

// DeleteLeave ...
func (LeaveService) DeleteLeave(id string) (e error) {
	defer config.CatchError(&e)
	return leaveDao.DeleteLeave(id)
}
