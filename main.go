package main

import (
	"encoding/json"
	"fmt"
	"time"
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
	Fit(int) bool
}

type BedStatus int

const (
	Unoccupied BedStatus = iota
	Reserved
	Occupied
)

type ShelterStatus struct {
	Beds        []BedStatus
	Occupants   []User
	LastUpdated time.Time
}

type ShelterCredentials struct {
	Username string
	Password string
}

type OpenTime struct {
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
	OpenTimes map[DayOfTheWeek][]OpenTime
}

type RequirementImportance int

const (
	UnImportant RequirementImportance = iota
	VeryImportant
	Necessary
)

type RequirmentMeta struct {
	Name       ShelterRequirementName
	Importance RequirementImportance
}

type CategoricalRequirement struct {
	Expected []int
	RequirmentMeta
}

func (req CategoricalRequirement) Fit(actual int) bool {
	for _, expected := range req.Expected {
		if actual == expected {
			return true
		}
	}
	return false
}

// RangeRequirement defines reqirment that x be in range [a, b)
type RangeRequirement struct {
	Expected [][2]int
	RequirmentMeta
}

func (req RangeRequirement) Fit(actual int) bool {
	for _, expected := range req.Expected {
		lowerBound := expected[0]
		upperBound := expected[1]
		if actual >= lowerBound && actual < upperBound {
			return true
		}
	}
	return false
}

type Shelter struct {
	Id              int // Database ID
	Name            string
	Requirements    []ShelterRequirement
	CurrentStatus   ShelterStatus
	Credentials     []ShelterCredentials
	Archive         []ShelterStatus
	Location        Location
	ShelterSchedule ShelterSchedule
}

type Location struct {
	Address string
	Lat     float64
	Lng     float64
}

type UserInfo struct {
	Id          int // Database ID
	FirstName   string
	MiddleName  string
	LastName    string
	PhoneNumber string
	Address     Location
}

type Boolean int

const (
	Unknown Boolean = iota
	No
	Yes
	Maybe
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
	WeeksPregnent   int
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
	Location          Location
	ReservationStatus ReservationStatus
	Shelter           Shelter
	IsSober           Boolean
	LastUpdated       time.Time
}

type User struct {
	// Maps relation type to user info
	Relationships  map[string]UserInfo
	IsUser         int // Determines whether this is a user. if missing is a UserInfo
	CurrentStatus  UserStatus
	ArchivedStatus []UserStatus
	PersonalInfo   PersonalInfo
	UserInfo
}

type Party []User

type Error struct {
	Number       uint
	ResponseCode uint
	ErrString    string
}

func (err *Error) Error() string {
	return err.ErrString
}

type CheckInForm struct {
}

func main() {
	req := RangeRequirement{
		RequirmentMeta: RequirmentMeta{
			Name: AgeReq,
		},
		Expected: [][2]int{
			[2]int{0, 100},
		},
	}
	user := User{}
	shelter := Shelter{Requirements: []ShelterRequirement{req}}
	thing, _ := json.MarshalIndent(user, "", " ")
	thing2, _ := json.MarshalIndent(shelter, "", " ")
	fmt.Println(string(thing))
	fmt.Println(string(thing2))

}
