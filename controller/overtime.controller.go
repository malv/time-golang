package controller

import (
	"time-project/config"
	"time-project/model"
	"time-project/service"

	"github.com/labstack/echo"
)

var overtimeService service.OvertimeService = service.OvertimeService{}

// SetOvertimeController ...
func SetOvertimeController(e *echo.Group) {
	e.GET("/overtimes", getOvertimes)
	e.GET("/overtime/:id", getOvertimeByID)
	e.POST("/overtimes", addOvertimes)
	e.PUT("/overtime", editOvertimes)
	e.DELETE("/overtime/:id", deleteOvertimes)
}

func getOvertimes(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := overtimeService.GetOvertimes()

	if err != nil {
		return resErr(c, err)
	}

	if len(result) == 0 {
		return res(c, []model.Overtime{})
	}
	return res(c, result)
}

func getOvertimeByID(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := overtimeService.GetOvertimeByID(id)

	if err != nil {
		return resErr(c, err)
	}
	return res(c, result)
}

func addOvertimes(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.Overtime{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := overtimeService.AddOvertime(data)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func editOvertimes(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.Overtime{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := overtimeService.EditOvertime(data)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func deleteOvertimes(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	err := overtimeService.DeleteOvertime(id)

	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}
