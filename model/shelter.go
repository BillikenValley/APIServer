package model

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

	MaxMaleAge         ShelterRequirementName = "max_male_age"
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

type ShelterCredentials struct {
	gorm.Model

	Username string
	Password string
}

type ShelterSchedule struct {
	gorm.Model

	OpenTime    time.Time `json:"open_time"`
	ClosingTime time.Time `json:"closing_time"`
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

type RequirementImportance int

const (
	UnImportant RequirementImportance = iota
	VeryImportant
	Necessary
)

type Requirement struct {
	gorm.Model

	Name
}

// RangeRequirement defines reqirment that x be in range [a, b)
type ShelterConstraints struct {
	gorm.Model

	AcceptsMen
	AcceptsSingleMen
	MaxMaleAge
	MinMaleAge
	AcceptsWomen
}

type ShelterID int

type Shelter struct {
	gorm.Model

	ShelterID       ShelterID `json:"uuid"`
	Name            string    `json:"name"`
	Requirements    map[string]Requirement
	CurrentStatus   ShelterStatus
	Credentials     []ShelterCredentials
	Archive         []ShelterStatus
	Location        Location
	ShelterSchedule ShelterSchedule
}

var shelters map[ShelterID]Shelter

func ShelterIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(shelters)
}

func ShelterShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shelterIDStr := vars["ShelterID"]
	val, _ := strconv.Atoi(shelterIDStr)
	var shelterID ShelterID = ShelterID(val)
	if shelter, ok := shelters[shelterID]; ok {
		json.NewEncoder(w).Encode(shelter)
	} else {
		json.NewEncoder(w).Encode(NewError(BadId, 404, fmt.Errorf("Bad ID Field")))
	}
	json.NewEncoder(w)
}

func ShelterUpload(w http.ResponseWriter, r *http.Request) {
	var shelter Shelter
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MAX_UPLOAD_SIZE))
	if err != nil {
		json.NewEncoder(w).Encode(NewError(BadRequestBody, 400, fmt.Errorf("BadRequestBody")))
	}
	if err := r.Body.Close(); err != nil {
		json.NewEncoder(w).Encode(NewError(ServerErr, 500, err))
	}
	if err := json.Unmarshal(body, &shelter); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	shelter.ShelterID = ShelterID(rand.Int())
	shelters[shelter.ShelterID] = shelter
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shelter)
}
