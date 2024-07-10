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
	return 1, nil
}

// InsertRoomRestriction inserts a room restrictions into the database
func (m *testDbRepo) InsertRoomRestriction(r models.RoomRestriction) error {
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
