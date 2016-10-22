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

// Valid Reqirments
const (
	AgeReq             = "age"
	SexReq             = "sex"
	GenderReq          = "gender"
	IsPregnant         = "is_pregnant"
	WeeksPregnant      = "weeks_pregnant"
	IsSexOffender      = "is_sex_offender"
	IsViolentCriminal  = "is_violent_criminal"
	HasId              = "has_id"
	HasSSN             = "has_ssn"
	IsWorking          = "is_working"
	InSchool           = "in_school"
	IsSober            = "is_sober"
	CurrentShelterUUID = "current_shelter_uuid"

	MaxMaleAge         = "max_male_age"
	MinMaleAge         = "max_male_age"
	MaxFemaleAge       = "max_female_age"
	MinFemaleAge       = "min_female_age"
	Children           = "num_children"
	TransFriendly      = "trans_friendly"
	AcceptsMen         = "accepts_men"
	AcceptsSingleMen   = "accepts_single_men"
	AcceptsWomen       = "accepts_women"
	AcceptsSingleWomen = "accepts_single_women"
	AcceptsChildren    = "accepts_children"
	MaxChildren        = "max_children"
	MustBeSober        = "must_be_sober"
)

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

type RequirementImportance int

const (
	UnImportant RequirementImportance = iota
	VeryImportant
	Necessary
)

type Requirement struct {
	Importance RequirementImportance
	Value      int
}

type ShelterID int

type Shelter struct {
	gorm.Model

	ShelterID          ShelterID              `json:"uuid"`
	Name               string                 `json:"name"`
	Beds               []User                 `json:"beds"`
	ShelterConstraints map[string]Requirement `json:"constraints"`
	ShelterSchedule    ShelterSchedule        `json:"schedule"`
	ShelterCredentials ShelterCredentials     `json:"-"`
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
