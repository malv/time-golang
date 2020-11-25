package main

import (
	"time-project/config"
	"time-project/controller"
	"time-project/dao"
	"time-project/service"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func main() {
	echo := echo.New()
	g := initDb()
	e := echo.Group("/api")

	dao.SetDao(g)
	service.SetService(g)

	controller.SetInit(e)
	controller.SetLeaveController(e)
	controller.SetPermitController(e)
	controller.SetOvertimeController(e)
	controller.SetLeaveRequestController(e)
	echo.Logger.Fatal(echo.Start(":9999"))
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		config.MigrateSchema(g)
		return g
	}
	panic(err)
}
