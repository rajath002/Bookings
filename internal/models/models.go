package models

import "time"

// Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Users in the user model.
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Rooms in the room model.
type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restrictions in the restriction model.
type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservations in the reservation model.
type Reservations struct {
	ID        int
	FirstName string
	LastName  string
	email     string
	Phone     string
	startDate time.Time
	endDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Rooms
}

// RoomRestrictions in the roomRestriction model.
type RoomRestrictions struct {
	ID             int
	startDate      time.Time
	endDate        time.Time
	RoomID         int
	ReservationID  int
	RestricationID int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Room           Rooms
	Reservation    Reservations
	Restriction    Restrictions
}
