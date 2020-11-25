package service

import "gorm.io/gorm"

var g *gorm.DB

// SetService ...
func SetService(gDB *gorm.DB) {
	g = gDB
}
