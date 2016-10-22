package model

import "github.com/jinzhu/gorm"

type Location struct {
	gorm.Model

	Address string
	Lat     float64
	Lng     float64
}

type Boolean int

const (
	Unknown Boolean = iota
	No
	Yes
	Maybe
)

type Match struct {
	Party    []User
	Shelters []Shelter
}
