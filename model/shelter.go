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
	AgeReq                 = "age"
	SexReq                 = "sex"
	GenderReq              = "gender"
	MustBePregnant         = "is_pregnant"
	WeeksPregnant          = "weeks_pregnant"
	AcceptsSexOffenders    = "is_sex_offender"
	AcceptsViolentCriminal = "is_violent_criminal"
	MustHaveID             = "has_id"
	MustHaveSSN            = "has_ssn"
	IsWorking              = "is_working"
	InSchool               = "in_school"
	AcceptsDrunks          = "is_sober"
	CurrentShelterUUID     = "current_shelter_uuid"

	MaxMaleAge         = "max_male_age"
	MinMaleAge         = "min_male_age"
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

type Shelter struct {
	gorm.Model

	ShelterID          int                `json:"uuid"`
	Name               string             `json:"name"`
	Beds               []int              `json:"beds"`
	Location           Location           `json:"location"`
	ShelterConstraints map[string]int     `json:"constraints"`
	OpenTime           time.Time          `json:"open_time"`
	ClosingTime        time.Time          `json:"close_time"`
	ShelterCredentials ShelterCredentials `json:"-"`
}

func RandomShelter() Shelter {
	beds := make([]int, 5)
	shelter := Shelter{
		ShelterID: rand.Int(),
		Name:      randSeq(5),
		Beds:      beds,
		Location: Location{
			Address: randSeq(5),
			Lat:     rand.Float64(),
			Lng:     rand.Float64(),
		},
		ShelterConstraints: map[string]int{
			MaxMaleAge:   rand.Int() % 100,
			MinMaleAge:   0,
			MaxFemaleAge: rand.Int() % 100,
			MinFemaleAge: 0,

			AcceptsSingleMen:    rand.Int() % 2,
			AcceptsSingleWomen:  rand.Int() % 2,
			AcceptsChildren:     rand.Int() % 2,
			AcceptsDrunks:       rand.Int() % 2,
			MustHaveID:          rand.Int() % 2,
			MustHaveSSN:         rand.Int() % 2,
			AcceptsSexOffenders: rand.Int() % 2,
		},
		OpenTime:    time.Now(),
		ClosingTime: time.Now().Add(2 * time.Hour),
		ShelterCredentials: ShelterCredentials{
			Username: "david",
			Password: "password",
		},
	}
	return shelter
}

func ShelterIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(shelters)
}

func ShelterShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shelterIDStr := vars["int"]
	val, _ := strconv.Atoi(shelterIDStr)
	var shelterID int = int(val)
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
	shelter.ShelterID = int(rand.Int())
	shelters[shelter.ShelterID] = shelter
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shelter)
}
