package config

import (
	"fmt"
	"time-project/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var tables = []interface{}{
	&model.Leave{},
	&model.Overtime{},
	&model.Permit{},
	&model.LeaveRequest{},
	&model.OvertimeRequest{},
	&model.PermitRequest{},
}

const (
	host     = "103.30.180.34"
	port     = 9595
	user     = "postgres"
	password = "bootlawen123"
	dbname   = "time"
	sslmode  = "disable"
)

// Conn ...
func Conn() (*gorm.DB, error) {
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// MigrateSchema ...
func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
