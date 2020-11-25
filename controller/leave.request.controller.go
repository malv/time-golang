package controller

import (
	"time-project/config"
	"time-project/model"
	"time-project/service"

	"github.com/labstack/echo"
)

var leaveRequestService service.LeaveRequestService = service.LeaveRequestService{}

// SetLeaveRequestController ...
func SetLeaveRequestController(e *echo.Group) {
	e.GET("/leave-requests", getLeaveRequests)
	e.GET("/leave-request/:id", getLeaveRequestByID)
	e.POST("/leave-requests", addLeaveRequest)
	e.PUT("/leave-requests", editLeaveRequest)
	e.DELETE("/leave-request/:id", deleteLeaveRequest)
}

func getLeaveRequests(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := leaveRequestService.GetLeaveRequests()

	if err != nil {
		return resErr(c, err)
	}

	if len(result) == 0 {
		return res(c, []model.LeaveRequest{})
	}
	return res(c, result)
}

func getLeaveRequestByID(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := leaveRequestService.GetLeaveRequestByID(id)

	if err != nil {
		resErr(c, err)
	}
	return res(c, result)
}

func addLeaveRequest(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.LeaveRequest{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := leaveRequestService.AddLeaveRequest(data)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func editLeaveRequest(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.LeaveRequest{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := leaveRequestService.EditLeaveRequest(data)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func deleteLeaveRequest(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	err := leaveRequestService.DeleteLeaveRequest(id)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}
