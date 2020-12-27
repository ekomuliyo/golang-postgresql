package models

import (
	"fmt"
	"golang-postgresql/db"
	"net/http"

	validator "github.com/go-playground/validator"
)

// Room ..
type Room struct {
	IDRoom      int    `json:"id_room,omitempty"`
	NameRoom    string `json:"name_room" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	MaxCapacity int    `json:"max_capacity" validate:"required"`
}

// StoreRoom ..
func StoreRoom(nameRoom string, price int, maxCapacity int) (Response, error) {

	var res Response

	v := validator.New()
	room := Room{
		NameRoom:    nameRoom,
		Price:       price,
		MaxCapacity: maxCapacity,
	}

	e := v.Struct(room)
	if e != nil {
		return res, e
	}

	con := db.CreateCon()
	sqlStatement := "INSERT INTO rooms (name_room, price, max_capacity) VALUES ($1, $2, $3) RETURNING id_room"

	var idRoom int64
	err := con.QueryRow(sqlStatement, nameRoom, price, maxCapacity).Scan(&idRoom)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"id_room": idRoom,
	}

	return res, nil

}

// GetAllRoom ..
func GetAllRoom() (Response, error) {
	var room Room
	var arrayRoom []Room
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT id_room, name_room, price, max_capacity FROM rooms ORDER BY id_room"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&room.IDRoom, &room.NameRoom, &room.Price, &room.MaxCapacity)
		if err != nil {
			return res, err
		}

		arrayRoom = append(arrayRoom, room)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayRoom

	return res, nil
}

// DeleteRoom ..
func DeleteRoom(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM rooms WHERE id_room = $1"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

// UpdateRoom ..
func UpdateRoom(id int, nameRoom string, price int, maxCapacity int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE rooms SET name_room = $1, price = $2, max_capacity = $3 WHERE id_room = $4"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nameRoom, price, maxCapacity, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}
