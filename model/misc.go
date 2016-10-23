package model

import "github.com/jinzhu/gorm"

type Location struct {
	gorm.Model

	Address string
	Lat     float64
	Lng     float64
}
