package models

import (
	"database/sql"
	"fmt"
	"golang-postgresql/db"
	"golang-postgresql/helpers"
	"net/http"

	validator "github.com/go-playground/validator"
)

// User ..
type User struct {
	IDUser   int    `json:"id_user,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	IDGroup  int    `json:"id_group" validate:"required"`
}

// RegisterUser ..
func RegisterUser(username string, password string, email string, IDGroup int) (Response, error) {

	var res Response

	v := validator.New()

	user := User{
		Username: username,
		Password: password,
		Email:    email,
		IDGroup:  IDGroup,
	}

	e := v.Struct(user)

	if e != nil {
		return res, e
	}

	con := db.CreateCon()

	sqlStatement := "INSERT INTO users (username, password, email, id_group) VALUES ($1, $2, $3, $4) RETURNING id_user"

	var idUser int64
	err := con.QueryRow(sqlStatement, username, password, email, IDGroup).Scan(&idUser)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"id_user": idUser,
	}

	return res, nil

}

// LoginUser ..
func LoginUser(email string, password string) (bool, int, error) {

	var user User
	var passwordHash string
	var IDGroup int

	con := db.CreateCon()

	sqlStatement := "SELECT email, password, id_group FROM users WHERE email = $1"

	err := con.QueryRow(sqlStatement, email).Scan(
		&user.Email, &passwordHash, &IDGroup,
	)

	if err == sql.ErrNoRows {
		fmt.Println("User not found")
		return false, 0, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, 0, err
	}

	matchPassword, err := helpers.CheckPasswordHash(password, passwordHash)

	if !matchPassword {
		fmt.Println("Hash and password doesn't match")
		return false, 0, err
	}

	return true, IDGroup, nil
}

// GetAllUser ..
func GetAllUser() (Response, error) {
	var user User
	var arrayUser []User
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT id_user, username, email, id_group FROM users ORDER BY id_user"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&user.IDUser, &user.Username, &user.Email, &user.IDGroup)
		if err != nil {
			return res, err
		}

		arrayUser = append(arrayUser, user)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayUser

	return res, nil
}

// UpdateUser ..
func UpdateUser(id int, username string, email string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE users SET username = $1, email = $2 WHERE id_user = $3"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, email, id)

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

// DeleteUser ..
func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM users WHERE id_user = $1"

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
