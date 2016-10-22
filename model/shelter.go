package model

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type ShelterRequirementName string

// Valid Reqirments
const (
	AgeReq          ShelterRequirementName = "age"
	SexReq          ShelterRequirementName = "sex"
	GenderReq       ShelterRequirementName = "gender"
	IsPregnant      ShelterRequirementName = "is_pregnant"
	WeeksPregnant   ShelterRequirementName = "weeks_pregnant"
	SexOffender     ShelterRequirementName = "sex_offender"
	ViolentCriminal ShelterRequirementName = "violent_criminal"
	HasId           ShelterRequirementName = "has_id"
	HasSSN          ShelterRequirementName = "has_ssn"
	IsWorking       ShelterRequirementName = "is_working"
	InSchool        ShelterRequirementName = "in_school"
	IsSober         ShelterRequirementName = "is_sober"

	MaleAge            ShelterRequirementName = "male_age"
	FemaleAge          ShelterRequirementName = "female_age"
	Children           ShelterRequirementName = "num_children"
	TransFriendly      ShelterRequirementName = "trans_friendly"
	AcceptsSingleMen   ShelterRequirementName = "accepts_single_men"
	AcceptsSingleWomen ShelterRequirementName = "accepts_single_women"
	MustBeSober        ShelterRequirementName = "must_be_sober"
)

type ShelterRequirement interface {
}

type BedStatus int

const (
	Unoccupied BedStatus = iota
	Reserved
	Occupied
)

type ShelterStatusID int

type ShelterStatus struct {
	gorm.Model

	ShelterStatusID int
	Beds            []BedStatus
	Occupants       []User
	LastUpdated     time.Time
}

type ShelterCredentials struct {
	gorm.Model

	Username string
	Password string
}

type OpenTime struct {
	gorm.Model

	From time.Time
	To   time.Time
}

type DayOfTheWeek int

const (
	Sun DayOfTheWeek = iota
	Mon
	Tues
	Wed
	Thurs
	Fri
	Sat
)

type ShelterSchedule struct {
	gorm.Model
	OpenTimes [7][]OpenTime
}

type RequirementImportance int

const (
	UnImportant RequirementImportance = iota
	VeryImportant
	Necessary
)

// RangeRequirement defines reqirment that x be in range [a, b)
type RangeRequirement struct {
	gorm.Model

	Name       ShelterRequirementName
	Importance RequirementImportance
	Expected   [][2]int
}

type Shelter struct {
	gorm.Model

	ShelterID       int // Database ID
	Name            string
	Requirements    []RangeRequirement
	CurrentStatus   ShelterStatus
	Credentials     []ShelterCredentials
	Archive         []ShelterStatus
	Location        Location
	ShelterSchedule ShelterSchedule
}

var shelters []Shelter
func ShelterIndex(w http.ResponseWriter, r *http.Request) {
	shelterJSON, err :=
}
