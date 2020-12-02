package controller

import (
	"time-project/config"
	"time-project/service"

	"github.com/labstack/echo"
)

var employeeService service.EmployeeService

// SetEmployeeController ...
func SetEmployeeController(e *echo.Group) {
	e.GET("/employees", getEmployees)
}

func getEmployees(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := employeeService.GetEmployees()

	if err != nil {
		return resErr(c, err)
	}

	return res(c, result)
}
