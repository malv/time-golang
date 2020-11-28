package controller

import (
	"time-project/config"
	"time-project/model"
	"time-project/service"

	"github.com/labstack/echo"
)

var overtimeRequestService service.OvertimeRequestService

// SetOvertimeRequestController ...
func SetOvertimeRequestController(e *echo.Group) {
	e.GET("/overtime-requests", getOvertimeRequests)
	e.GET("/overtime-request/:id", getOvertimeRequestByID)
	e.POST("/overtime-requests", addOvertimeRequest)
	e.PUT("/overtime-requests", editOvertimeRequest)
	e.DELETE("/overtime-request/:id", deleteOvertimeRequest)
}

func getOvertimeRequests(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := overtimeRequestService.GetOvertimeRequests()

	if err != nil {
		return resErr(c, err)
	}

	if len(result) == 0 {
		return res(c, []model.OvertimeRequest{})
	}
	return res(c, result)
}

func getOvertimeRequestByID(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := overtimeRequestService.GetOvertimeRequestByID(id)

	if err != nil {
		resErr(c, err)
	}
	return res(c, result)
}

func addOvertimeRequest(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.OvertimeRequest{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := overtimeRequestService.AddOvertimeRequest(data)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func editOvertimeRequest(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.OvertimeRequest{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := overtimeRequestService.EditOvertimeRequest(data)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func deleteOvertimeRequest(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	err := overtimeRequestService.DeleteOvertimeRequest(id)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}
