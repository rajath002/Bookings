package dbrepo

import (
	"errors"
	"time"

	"github.com/rajath002/bookings/internal/models"
)

func (m *testDbRepo) AllUsers() bool {
	return true
}

// Insert a reservation into the database
func (m *testDbRepo) InsertReservation(res models.Reservation) (int, error) {
	// if the room id is 2 then fail; otherwiser pass
	if res.RoomID == 2 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restrictions into the database
func (m *testDbRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("some error")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID, and false if no avaialbility.
func (m *testDbRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns Slice of available rooms for any given date Range.
func (m *testDbRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets a room by ID
func (m *testDbRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}
	return room, nil
}

// Returns a userByID
func (m *testDbRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

// UpdateUser updates a user in the database
func (m *testDbRepo) UpdateUser(u models.User) error {
	return nil
}

// Authenticate Authenticates a user
func (m *testDbRepo) Authenticate(email, testPassword string) (int, string, error) {
	if email == "me@here.ka" {
		return 1, "", nil
	}
	return 0, "", errors.New("some error")
}

// AllReservations returns a slice of all Reservations
func (m *testDbRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

// AllNewReservations returns a slice of all Reservations
func (m *testDbRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

func (m *testDbRepo) GetReservationByID(id int) (models.Reservation, error) {
	var reservation models.Reservation
	return reservation, nil
}

// UpdateReservation updates a reservation in the database
func (m *testDbRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

// UpdateProcessedForReservation updates processed for a reservation by id
func (m *testDbRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

// DeleteReservation a reservation by id
func (m *testDbRepo) DeleteReservation(id int) error {
	return nil
}

func (m *testDbRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRestricationsForRoomByDate returns restrictions for a room byt date range
func (m *testDbRepo) GetRestricationsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	var restrictions []models.RoomRestriction
	return restrictions, nil
}

// InsertBlockFromRoom Inserts a room restriction
func (m *testDbRepo) InsertBlockFromRoom(id int, startDate time.Time) error {
	return nil
}

// DeleteBlockById Deletes room restriction
func (m *testDbRepo) DeleteBlockById(id int) error {
	return nil
}
