package controller

import (
	"time-project/config"
	"time-project/model"
	"time-project/service"

	"github.com/labstack/echo"
)

var permitRequestService service.PermitRequestService

// SetPermitRequestController ...
func SetPermitRequestController(e *echo.Group) {
	e.GET("/permit-requests", getPermitRequests)
	e.GET("/permit-request/:id", getPermitRequestByID)
	e.POST("/permit-requests", addPermitRequest)
	e.PUT("/permit-requests", editPermitRequest)
	e.DELETE("/permit-request/:id", deletePermitRequest)
}

func getPermitRequests(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := permitRequestService.GetPermitRequests()

	if err != nil {
		return resErr(c, err)
	}

	if len(result) == 0 {
		return res(c, []model.PermitRequest{})
	}
	return res(c, result)
}

func getPermitRequestByID(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := permitRequestService.GetPermitRequestByID(id)

	if err != nil {
		resErr(c, err)
	}
	return res(c, result)
}

func addPermitRequest(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.PermitRequest{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := permitRequestService.AddPermitRequest(data)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func editPermitRequest(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.PermitRequest{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := permitRequestService.EditPermitRequest(data)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func deletePermitRequest(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	err := permitRequestService.DeletePermitRequest(id)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}
