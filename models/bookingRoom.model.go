package models

import (
	"golang-postgresql/db"
	"net/http"
)

type BookingRoom struct {
	IDBookingRoom   interface{} `json:"id_booking_room,omitempty"`
	IDRoom          int         `json:"id_room"`
	NameRoom        string      `json:"name_room"`
	Price           int         `json:"price"`
	MaxCapacity     int         `json:"max_capacity"`
	Username        interface{} `json:"username,omitempty"`
	DateBooking     interface{} `json:"date_booking,omitempty"`
	AmounCapacity   interface{} `json:"amoun_capacity,omitempty"`
	IDStatusBooking int         `json:"id_status_booking"`
}

func GetAllBookingRoom() (Response, error) {

	var bookingRoom BookingRoom
	var arrayBookingRoom []BookingRoom
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT b.id_booking_room, a.id_room, a.name_room, a.price, a.max_capacity, c.username, 
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
			&bookingRoom.Username, &bookingRoom.DateBooking, &bookingRoom.AmounCapacity, &bookingRoom.IDStatusBooking)
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
