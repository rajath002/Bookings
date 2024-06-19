package dbrepo

import (
	"time"

	"github.com/rajath002/bookings/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// Insert a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction inserts a room restrictions into the database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID, and false if no avaialbility.
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns Slice of available rooms for any given date Range.
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomID gets a room by ID
func (m *postgresDBRepo) GetRoomID(id int) (models.Room, error) {
	var room models.Room
	return room, nil
}
