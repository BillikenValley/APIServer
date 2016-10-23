package model

import "math/rand"

type Gender int

const (
	GenderUnknown Gender = iota
	Male
	Female
	OtherGender
)

type Requester struct {
	Age               int    `json:"age"`
	Gender            Gender `json:"gender"`
	Sex               Gender `json:"sex"`
	IsVeteran         int    `json:"is_veteran"`
	IsPregnant        int    `json:"is_pregnant"`
	WeeksPregnant     int    `json:"weeks_pregnant"`
	IsSexOffender     int    `json:"is_sex_offender"`
	IsViolentCriminal int    `json:"is_violent_criminal"`
	CurrentShelter    int    `json:"current_shelter_uuid"`
}

func RandomRequester() Requester {
	return Requester{
		Age:               rand.Int() % 100,
		Gender:            Gender(rand.Int() % 3),
		Sex:               Gender(rand.Int() % 3),
		IsVeteran:         rand.Int() % 2,
		WeeksPregnant:     rand.Int() % 12,
		IsSexOffender:     rand.Int() % 2,
		IsViolentCriminal: rand.Int() % 2,
	}
}
