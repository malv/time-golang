package dao

import (
	"time-project/config"
	"time-project/model"
)

// PermitRequestDao ...
type PermitRequestDao struct{}

// GetPermitRequests ...
func (PermitRequestDao) GetPermitRequests() (data []model.PermitRequest, e error) {
	defer config.CatchError(&e)

	var permitRequest []model.PermitRequest

	result := g.Find(&permitRequest)

	if result.Error != nil {
		return data, result.Error
	}
	return permitRequest, nil
}

// GetPermitRequestByID ...
func (PermitRequestDao) GetPermitRequestByID(id string) (data model.PermitRequest, e error) {
	defer config.CatchError(&e)

	permitRequest := model.PermitRequest{BaseModel: model.BaseModel{ID: id}}

	result := g.First(&permitRequest)

	if result.Error != nil {
		return data, result.Error
	}
	return permitRequest, nil
}

// AddPermitRequest ...
func (PermitRequestDao) AddPermitRequest(data *model.PermitRequest) (e error) {
	defer config.CatchError(&e)

	data.CreatedBy = data.EmployeeID
	data.UpdatedBy = &data.EmployeeID
	data.UpdatedAt = data.CreatedAt
	data.Status = "P"

	result := g.Create(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// EditPermitRequest ...
func (PermitRequestDao) EditPermitRequest(data *model.PermitRequest) (e error) {
	defer config.CatchError(&e)

	result := g.Save(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeletePermitRequest ...
func (PermitRequestDao) DeletePermitRequest(id string) (e error) {
	defer config.CatchError(&e)

	var permitRequest model.PermitRequest

	result := g.Delete(&permitRequest, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
