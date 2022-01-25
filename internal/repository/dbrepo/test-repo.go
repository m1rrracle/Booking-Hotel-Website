package dbrepo

import (
	"errors"
	"pdev-finmart/internal/models"
	"time"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

//InsertRoomRestriction insert a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error  {
	return nil
}

// SearchAvailabilityByDates return true if availability exist for roomID, and false if no availability
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)  {
	return false, nil
}

//SearchAvailabilityForAllRooms return a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)  {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets a room by ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error)  {
	var room models.Room
	return room, nil
}

// GetUserByID returns a user by ID
func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User
	return u, nil
}

//UpdateUser update a user in the database
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

//Authenticate authenticate a user
func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error)  {
	if email == "admin@magang.com" {
		return 1, "", nil
	}
	return 0, "", errors.New("some error")
}

// AllReservations returns a slice of all reservations
func (m *testDBRepo) AllReservations() ([]models.Reservation, error)  {
	var reservations []models.Reservation
	return reservations, nil
}

// AllNewReservations returns a slice of all reservations
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error)  {
	var reservations []models.Reservation
	return reservations, nil
}

// GetReservationByID return one reservation by ID
func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation
	return res, nil
}

func (m *testDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

func (m *testDBRepo) DeleteReservation(id int) error  {
	return nil
}

func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error  {
	return nil
}

func (m *testDBRepo) AllRooms()([]models.Room, error)  {
	var rooms []models.Room
	return rooms, nil
}
func (m *testDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error)  {
	var restrictions []models.RoomRestriction
	return restrictions, nil
}

func (m *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error  {
	return nil
}

func (m *testDBRepo) DeleteBlockByID(id int) error  {
	return nil
}