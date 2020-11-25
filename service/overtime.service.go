package service

import (
	"time-project/config"
	"time-project/dao"
	"time-project/model"
)

// OvertimeService ...
type OvertimeService struct{}

var overtimeDao dao.OvertimeDao = dao.OvertimeDao{}

// GetOvertimes ...
func (OvertimeService) GetOvertimes() (data []model.Overtime, e error) {
	defer config.CatchError(&e)

	result, err := overtimeDao.GetOvertimes()

	if err != nil {
		return data, err
	}

	if len(result) == 0 {
		return []model.Overtime{}, nil
	}
	return result, nil
}

// GetOvertimeByID ...
func (OvertimeService) GetOvertimeByID(id string) (data model.Overtime, e error) {
	defer config.CatchError(&e)

	result, err := overtimeDao.GetOvertimeByID(id)

	if err != nil {
		return data, err
	}
	return result, nil
}

// AddOvertime ...
func (OvertimeService) AddOvertime(data *model.Overtime) (e error) {
	defer config.CatchError(&e)
	data.CreatedBy = "Admin"
	admin := "Admin"
	data.UpdatedBy = &admin
	return overtimeDao.AddOvertime(data)
}

// EditOvertime ...
func (OvertimeService) EditOvertime(data *model.Overtime) (e error) {
	defer config.CatchError(&e)
	return overtimeDao.EditOvertime(data)
}

// DeleteOvertime ...
func (OvertimeService) DeleteOvertime(id string) (e error) {
	defer config.CatchError(&e)
	return overtimeDao.DeleteOvertime(id)
}
