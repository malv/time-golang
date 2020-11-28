package service

import (
	"time-project/config"
	"time-project/dao"
	"time-project/model"
)

// PermitRequestService ...
type PermitRequestService struct{}

var permitRequestDao dao.PermitRequestDao = dao.PermitRequestDao{}

// GetPermitRequests ...
func (PermitRequestService) GetPermitRequests() (data []model.PermitRequest, e error) {
	defer config.CatchError(&e)

	result, err := permitRequestDao.GetPermitRequests()

	if err != nil {
		return data, err
	}
	return result, nil
}

// GetPermitRequestByID ...
func (PermitRequestService) GetPermitRequestByID(id string) (data model.PermitRequest, e error) {
	defer config.CatchError(&e)

	result, err := permitRequestDao.GetPermitRequestByID(id)

	if err != nil {
		return data, err
	}
	return result, nil
}

// AddPermitRequest ...
func (PermitRequestService) AddPermitRequest(data *model.PermitRequest) (e error) {
	defer config.CatchError(&e)
	return permitRequestDao.AddPermitRequest(data)
}

// EditPermitRequest ...
func (PermitRequestService) EditPermitRequest(data *model.PermitRequest) (e error) {
	defer config.CatchError(&e)
	return permitRequestDao.EditPermitRequest(data)
}

// DeletePermitRequest ...
func (PermitRequestService) DeletePermitRequest(id string) (e error) {
	defer config.CatchError(&e)
	return permitRequestDao.DeletePermitRequest(id)
}
