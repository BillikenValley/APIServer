package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	gorm.Model

	UserID      int // Database ID
	FirstName   string
	MiddleName  string
	LastName    string
	PhoneNumber string
	Address     Location
}

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
	Relationships  map[string]UserInfo
	IsUser         Boolean // Determines whether this is a user. if missing is a UserInfo
	CurrentStatus  UserStatus
	ArchivedStatus []UserStatus
	PersonalInfo   PersonalInfo
	UserInfo
}
