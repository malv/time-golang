package controller

import (
	"time-project/config"
	"time-project/model"
	"time-project/service"

	"github.com/labstack/echo"
)

var leaveService service.LeaveService = service.LeaveService{}

// SetLeaveController ...
func SetLeaveController(e *echo.Group) {
	e.GET("/leaves", getLeaves)
	e.GET("/leave/:id", getLeaveByID)
	e.POST("/leaves", addLeave)
	e.PUT("/leaves", editLeave)
	e.DELETE("leaves", deleteLeave)
}

func getLeaves(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := leaveService.GetLeaves()

	if err != nil {
		return resErr(c, err)
	}

	if len(result) > 0 {
		return res(c, result)
	}
	return res(c, []model.Leave{})
}

func getLeaveByID(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := leaveService.GetLeaveByID(id)

	if err != nil {
		return resErr(c, err)
	}

	return res(c, result)
}

func addLeave(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Leave{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := leaveService.AddLeave(data)
	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func editLeave(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Leave{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := leaveService.EditLeave(data)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func deleteLeave(c echo.Context) (e error) {
	defer config.CatchError(&e)
	id := c.Param("id")

	err := leaveService.DeleteLeave(id)

	if err != nil {
		return resErr(c, err)
	}

	return resSuccess(c)
}
