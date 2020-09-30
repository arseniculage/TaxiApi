package TaxiApi

import (
	"database/sql/driver"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
	"image"
	"time"
)
const(
	Card = iota
	Cash
)
//User structure
type User struct {
	UserId uuid.UUID
	PhoneNumber [10]uint8
	Name string
	Email string
}
//Driver structure
type Driver struct{
	DriverId uuid.UUID
	PhoneNumber [10]uint8
	Name string
}
//Ticket structure
type Ticket struct {
	DriverId uuid.UUID
	UserId uuid.UUID
	ticketStartTime time.Time
	ticketFinishTime time.Time
	paymentType uint8
}

type APIServer struct{
	config *ServerConfig
	router *mux.Router
	db *Database
}
