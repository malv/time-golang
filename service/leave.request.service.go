package service

import (
	"time-project/config"
	"time-project/dao"
	"time-project/model"
)

// LeaveRequestService ...
type LeaveRequestService struct{}

var leaveRequestDao dao.LeaveRequestDao = dao.LeaveRequestDao{}

// GetLeaveRequests ...
func (LeaveRequestService) GetLeaveRequests() (data []model.LeaveRequest, e error) {
	defer config.CatchError(&e)

	result, err := leaveRequestDao.GetLeaveRequests()

	if err != nil {
		return data, err
	}
	return result, nil
}

// GetLeaveRequestByID ...
func (LeaveRequestService) GetLeaveRequestByID(id string) (data model.LeaveRequest, e error) {
	defer config.CatchError(&e)

	result, err := leaveRequestDao.GetLeaveRequestByID(id)

	if err != nil {
		return data, err
	}
	return result, nil
}

// AddLeaveRequest ...
func (LeaveRequestService) AddLeaveRequest(data *model.LeaveRequest) (e error) {
	defer config.CatchError(&e)
	return leaveRequestDao.AddLeaveRequest(data)
}

// EditLeaveRequest ...
func (LeaveRequestService) EditLeaveRequest(data *model.LeaveRequest) (e error) {
	defer config.CatchError(&e)
	return leaveRequestDao.EditLeaveRequest(data)
}

// DeleteLeaveRequest ...
func (LeaveRequestService) DeleteLeaveRequest(id string) (e error) {
	defer config.CatchError(&e)
	return leaveRequestDao.DeleteLeaveRequest(id)
}
