package model

import (
	"math/rand"

	"github.com/jinzhu/gorm"
)

type Guest struct {
	gorm.Model

	GuestID        int    `json:"id"`
	FirstName      string `json:"first_name"`
	MiddleName     string `json:"middle_name"`
	LastName       string `json:"last_name"`
	PhoneNumber    string `json:"phone_number"`
	CurrentShelter Shelter
}

func RandomGuest() Guest {
	return Guest{
		GuestID:     rand.Int(),
		FirstName:   randSeq(5),
		MiddleName:  randSeq(5),
		LastName:    randSeq(5),
		PhoneNumber: randSeq(7),
	}
}
