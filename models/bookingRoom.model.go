package models

import (
	"errors"
	"fmt"
	"golang-postgresql/db"
	"net/http"
	"strconv"

	validator "github.com/go-playground/validator"
)

type BookingRoom struct {
	IDBookingRoom   interface{} `json:"id_booking_room,omitempty"`
	IDRoom          int         `json:"id_room,omitempty" validate:"required"`
	NameRoom        string      `json:"name_room"`
	Price           int         `json:"price"`
	MaxCapacity     int         `json:"max_capacity"`
	Username        interface{} `json:"username,omitempty"`
	IDUser          interface{} `json:"id_user,omitempty" validate:"required"`
	DateBooking     interface{} `json:"date_booking,omitempty" validate:"required"`
	AmounCapacity   interface{} `json:"amoun_capacity,omitempty" validate:"required"`
	IDStatusBooking int         `json:"id_status_booking"`
}

func GetAllBookingRoom() (Response, error) {

	var bookingRoom BookingRoom
	var arrayBookingRoom []BookingRoom
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT b.id_booking_room, a.id_room, a.name_room, a.price, a.max_capacity, c.username, c.id_user,
							b.date_booking, b.amount_capacity, COALESCE(b.id_status_booking, 1) AS id_status_booking  
						FROM rooms a
					LEFT JOIN booking_room b on a.id_room = b.id_room
					LEFT JOIN users c on b.id_user = c.id_user`

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&bookingRoom.IDBookingRoom, &bookingRoom.IDRoom, &bookingRoom.NameRoom, &bookingRoom.Price, &bookingRoom.MaxCapacity,
			&bookingRoom.Username, &bookingRoom.IDUser, &bookingRoom.DateBooking, &bookingRoom.AmounCapacity, &bookingRoom.IDStatusBooking)
		if err != nil {
			return res, err
		}

		arrayBookingRoom = append(arrayBookingRoom, bookingRoom)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayBookingRoom

	return res, nil
}

func StoreBookingRoom(idRoom int, idUser int, dateBooking string, amount_capacity int) (Response, error) {

	var res Response

	v := validator.New()
	bookingRoom := BookingRoom{
		IDRoom:        idRoom,
		IDUser:        idUser,
		DateBooking:   dateBooking,
		AmounCapacity: amount_capacity,
	}

	e := v.Struct(bookingRoom)
	if e != nil {
		return res, e
	}

	con := db.CreateCon()

	// periksa apakah jumlah booking melebih kapasitas atau tidak
	sqlChecAmount := "SELECT max_capacity FROM rooms WHERE id_room = " + strconv.Itoa(idRoom)
	var maxCapacity int
	er := con.QueryRow(sqlChecAmount).Scan(&maxCapacity)
	if er != nil {
		fmt.Println(er)
		return res, er
	}

	if amount_capacity > maxCapacity {
		return res, errors.New("Jumlah Booking Anda Melebihi Kapasitas")
	}

	defer con.Close()

	sqlStatement := "INSERT INTO booking_room (id_room, id_user, date_booking, amount_capacity, id_status_booking) VALUES ($1, $2, $3, $4, $5) RETURNING id_booking_room"

	var idBookingRoom int64
	var idStatusBooking int = 2 /* Booked */
	err := con.QueryRow(sqlStatement, idRoom, idUser, dateBooking, amount_capacity, idStatusBooking).Scan(&idBookingRoom)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"id_booking_room": idBookingRoom,
	}

	return res, nil
}
