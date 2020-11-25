package dao

import (
	"time-project/config"

	"gorm.io/gorm"
)

var g *gorm.DB

// SetDao ...
func SetDao(gDB *gorm.DB) {
	g = gDB
}
func catchError(e *error) {
	config.CatchError(e)
}
