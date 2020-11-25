package controller

import (
	"time-project/config"
	"time-project/model"
	"time-project/service"

	"github.com/labstack/echo"
)

var permitService service.PermitService = service.PermitService{}

// SetPermitController ...
func SetPermitController(e *echo.Group) {
	e.GET("/permits", getPermits)
	e.GET("/permit/:id", getPermitByID)
	e.POST("/permits", addPermit)
	e.PUT("/permits", editPermit)
	e.DELETE("/permit/:id", deletePermit)
}

func getPermits(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := permitService.GetPermits()

	if err != nil {
		return resErr(c, err)
	}

	if len(result) == 0 {
		return res(c, []model.Permit{})
	}
	return res(c, result)
}

func getPermitByID(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := permitService.GetPermitByID(id)

	if err != nil {
		return resErr(c, err)
	}

	return res(c, result)
}

func addPermit(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.Permit{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := permitService.AddPermit(data)
	if err != nil {
		return resErr(c, err)
	}

	return resSuccess(c)
}

func editPermit(c echo.Context) (e error) {
	defer config.CatchError(&e)

	data := &model.Permit{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := permitService.EditPermit(data)

	if err != nil {
		return resErr(c, err)
	}

	return resSuccess(c)
}

func deletePermit(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")
	err := permitService.DeletePermit(id)

	if err != nil {
		return resErr(c, err)
	}

	return resSuccess(c)
}
