package service

import (
	"time-project/config"
	"time-project/dao"
	"time-project/model"
)

// PermitService ...
type PermitService struct{}

var permitDao dao.PermitDao = dao.PermitDao{}

// GetPermits ...
func (PermitService) GetPermits() (data []model.Permit, e error) {
	defer config.CatchError(&e)

	result, err := permitDao.GetPermits()

	if err != nil {
		return data, err
	}

	return result, nil
}

// GetPermitByID ...
func (PermitService) GetPermitByID(id string) (data model.Permit, e error) {
	defer config.CatchError(&e)

	result, err := permitDao.GetPermitByID(id)

	if err != nil {
		return data, err
	}

	return result, nil
}

// AddPermit ...
func (PermitService) AddPermit(data *model.Permit) (e error) {
	defer config.CatchError(&e)

	data.CreatedBy = "Admin"
	admin := "Admin"
	data.UpdatedBy = &admin
	return permitDao.AddPermit(data)
}

// EditPermit ...
func (PermitService) EditPermit(data *model.Permit) (e error) {
	defer config.CatchError(&e)

	return permitDao.EditPermit(data)
}

// DeletePermit ...
func (PermitService) DeletePermit(id string) (e error) {
	defer config.CatchError(&e)

	return permitDao.DeletePermit(id)
}
