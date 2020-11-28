package service

import (
	"time-project/config"
	"time-project/dao"
	"time-project/model"
)

// OvertimeRequestService ...
type OvertimeRequestService struct{}

var overtimeRequestDao dao.OvertimeRequestDao = dao.OvertimeRequestDao{}

// GetOvertimeRequests ...
func (OvertimeRequestService) GetOvertimeRequests() (data []model.OvertimeRequest, e error) {
	defer config.CatchError(&e)

	result, err := overtimeRequestDao.GetOvertimeRequests()

	if err != nil {
		return data, err
	}
	return result, nil
}

// GetOvertimeRequestByID ...
func (OvertimeRequestService) GetOvertimeRequestByID(id string) (data model.OvertimeRequest, e error) {
	defer config.CatchError(&e)

	result, err := overtimeRequestDao.GetOvertimeRequestByID(id)

	if err != nil {
		return data, err
	}
	return result, nil
}

// AddOvertimeRequest ...
func (OvertimeRequestService) AddOvertimeRequest(data *model.OvertimeRequest) (e error) {
	defer config.CatchError(&e)
	return overtimeRequestDao.AddOvertimeRequest(data)
}

// EditOvertimeRequest ...
func (OvertimeRequestService) EditOvertimeRequest(data *model.OvertimeRequest) (e error) {
	defer config.CatchError(&e)
	return overtimeRequestDao.EditOvertimeRequest(data)
}

// DeleteOvertimeRequest ...
func (OvertimeRequestService) DeleteOvertimeRequest(id string) (e error) {
	defer config.CatchError(&e)
	return overtimeRequestDao.DeleteOvertimeRequest(id)
}
