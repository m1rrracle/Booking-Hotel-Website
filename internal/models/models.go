package models

import (
	"time"
)

// User is the user model
type User struct {
	ID        int
	FullName  string
	Email     string
	Phone     string
	KTPImages       string
	Password  string
	IsActive  int
	AddressID string
	RoleID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      Role
	Addresses Addresses
}

// Role is the role model
type Role struct {
	ID          int
	Name    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Addresses is the addresses model
type Addresses struct {
	ID            int
	ProvinceID    int
	CityID        int
	DistrictID    int
	SubDistrictID int
	PostalCodeID  int
	Latitude      float64
	Longitude     float64
	Address       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Nation struct {
	NationCode    int
	Iso2          string
	Iso3          string
	Name          string
	Currency      string
	CurrCode      string
	CurrSymbol    string
	Weight        string
	Length        string
	Latitude      float64
	Longitude     float64
	IsProvince    int
	IsCity        int
	IsDistrict    int
	IsSubDistrict int
	IsPostalCode  int
	IsActive      int
}

type Province struct {
	ID         int
	NationCode int
	Latitude   float64
	Longitude  float64
	Nation     Nation
}

type City struct {
	ID         int
	ProvinceID int
	Name       string
	Latitude   float64
	Longitude  float64
	Province   Province
}

type District struct {
	ID        int
	CityID    int
	Name      string
	Latitude  float64
	Longitude float64
	City      City
}

type SubDistrict struct {
	ID         int
	DistrictID int
	Name       string
	Latitude   float64
	Longitude  float64
	District   District
}

type PostalCode struct {
	ID         int
	CityID     int
	DistrictID int
	PostalCode string
	City       City
	District   District
}

// Room is the room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restriction is the restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservation is the reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
	Processed int
}

// RoomRestriction is the room restriction model
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}

// MailData holds an email message
type MailData struct {
	To       string
	From     string
	Subject  string
	Content  string
	Template string
}
