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

type MilitaryBranch int

const (
	Army       MilitaryBranch = 1
	Navy       MilitaryBranch = 2
	Airforce   MilitaryBranch = 3
	Marines    MilitaryBranch = 4
	CoastGuard MilitaryBranch = 5
)

type Gender int

const (
	GenderUnknown Gender = iota
	Male
	Female
	OtherGender
)

type DischargeStatus int

// DataQuality 1 if this is guarenteed to be the name. 99 if there is no name.
// Booleans are in the fol
type PersonalInfo struct {
	gorm.Model

	PersonalInfoID int
	/* Information in the GlobalHack 6 Dataset */
	UUID       int
	FirstName  string
	MiddleName string
	LastName   string
	// NameDataQuality indiecates the quality of the name paramaters
	NameDataQuality int
	// 0 if not known
	SSN            int
	SSNDataQuality int
	// 0 if not known
	DOB            time.Time
	DOBDataQuality int

	AmIndAKNative        Boolean
	Asian                Boolean
	Black                Boolean
	NativeHIOtherPacific Boolean
	White                Boolean
	RaceNone             Boolean
	Gender               Gender
	OtherGender          Boolean
	VeteranStatus        Boolean

	YearEnteredService time.Time
	YearSeparated      time.Time
	WorldWarII         Boolean
	KoreanWar          Boolean
	VietnamWar         Boolean
	DesertStorm        Boolean
	AfghanistanOEF     Boolean
	IraqOIF            Boolean
	IraqOND            Boolean
	OtherTheater       Boolean
	MilitaryBranch     MilitaryBranch
	DischargeStatus    DischargeStatus

	DateCreated time.Time
	DateUpdated time.Time
	UserID      time.Time

	/* Information we added in our app */
	Pregnant        Boolean
	WeeksPregnant   int
	HIVStatus       Boolean
	Sex             Gender
	SexOffender     Boolean
	ViolentCriminal Boolean
	HasId           Boolean
	IsWorking       Boolean
	InSchool        Boolean
	MiscInfo        map[string]int
}

type ReservationStatus int

const (
	ReservationUnknown ReservationStatus = iota
	Unassigned
	OnRoute
	InShelter
	CheckedOut
)

type UserStatus struct {
	gorm.Model

	UserStatusID      int
	Location          Location
	ReservationStatus ReservationStatus
	ShelterId         int64
	IsSober           Boolean
	LastUpdated       time.Time
}

type UserID int64

type User struct {
	// Maps relation type to user info
	Relationships  map[string]User
	IsUser         Boolean // Determines whether this is a user. if missing is a UserInfo
	CurrentStatus  UserStatus
	ArchivedStatus []UserStatus
	PersonalInfo   PersonalInfo

	UserID      UserID // Database ID
	FirstName   string
	MiddleName  string
	LastName    string
	PhoneNumber string
	Address     Location
}

var users map[UserID]User

func UserIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["UserID"]
	val, _ := strconv.Atoi(userIDStr)
	var userID UserID = UserID(val)
	if user, ok := users[userID]; ok {
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode(NewError(BadId, 404, fmt.Errorf("Bad ID Field")))
	}
	json.NewEncoder(w)
}

func UserUpload(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MAX_UPLOAD_SIZE))
	if err != nil {
		json.NewEncoder(w).Encode(NewError(BadRequestBody, 400, fmt.Errorf("BadRequestBody")))
	}
	if err := r.Body.Close(); err != nil {
		json.NewEncoder(w).Encode(NewError(ServerErr, 500, err))
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	user.UserID = UserID(rand.Int())
	users[user.UserID] = user
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
