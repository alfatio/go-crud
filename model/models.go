package model

import (
	"errors"

	"github.com/alfatio/login/config"
	"github.com/alfatio/login/helper"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var db = config.DB()

func UserCol(colName string, user *User) interface{} {
	switch colName {
	case "user_id":
		return &user.Id
	case "username":
		return &user.Username
	case "password":
		return &user.Password
	case "email":
		return &user.Email
	default:
		panic("unkown column name " + colName)
	}
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users ORDER BY user_id"
	var output []User

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	colNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	colNum := len(colNames)

	defer rows.Close()

	for rows.Next() {
		var u User

		cols := make([]interface{}, colNum)

		for i := 0; i < colNum; i++ {
			cols[i] = UserCol(colNames[i], &u)

		}

		err := rows.Scan(cols...)

		if err != nil {
			return nil, err
		}

		output = append(output, u)
	}

	if len(output) == 0 {
		return output, errors.New("internal server error")
	}

	return output, nil

}

func GetUserByUsername(u string) (User, error) {
	var output User

	query := `
	SELECT * FROM users WHERE username = $1 LIMIT 1
	`
	rows, err := db.Query(query, u)
	if err != nil {
		return output, err
	}

	colNames, err := rows.Columns()
	if err != nil {
		return output, err
	}

	colNum := len(colNames)

	defer rows.Close()

	for rows.Next() {
		var u User

		cols := make([]interface{}, colNum)
		for i := 0; i < colNum; i++ {
			cols[i] = UserCol(colNames[i], &u)
		}

		err := rows.Scan(cols...)
		if err != nil {
			return output, err
		}

		output = u
	}

	if output.Id == 0 {
		return output, errors.New("no user with given username exist")
	}

	return output, nil

}

func InsertUser(p User) (User, error) {
	var output User
	h, err := helper.HashPW(p.Password)

	if err != nil {
		return output, err
	}

	p.Password = h

	query := `
	INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3)
	RETURNING *
	`
	rows, err := db.Query(query, p.Username, p.Password, p.Email)

	if err != nil {
		return output, err
	}

	defer rows.Close()

	colNames, err := rows.Columns()
	if err != nil {
		return output, err
	}

	colNum := len(colNames)

	for rows.Next() {
		var u User

		cols := make([]interface{}, colNum)

		for i := 0; i < colNum; i++ {
			cols[i] = UserCol(colNames[i], &u)
		}

		err := rows.Scan(cols...)
		if err != nil {
			return output, err
		}

		output = u
	}

	return output, nil
}

func EditUser(p User) (res User, err error) {

	var output User
	h, err := helper.HashPW(p.Password)

	if err != nil {
		return output, err
	}

	p.Password = h

	query := `
		UPDATE users
		SET username = $1,
				password = $2,
				email 	 = $3
		WHERE user_id = $4
    RETURNING *
	`
	rows, err := db.Query(query, p.Username, p.Password, p.Email, p.Id)
	if err != nil {
		return output, err
	}

	defer rows.Close()

	colNames, err := rows.Columns()
	if err != nil {
		return output, err
	}

	colNum := len(colNames)

	for rows.Next() {
		var u User

		cols := make([]interface{}, colNum)

		for i := 0; i < colNum; i++ {
			cols[i] = UserCol(colNames[i], &u)

		}

		err := rows.Scan(cols...)

		if err != nil {
			return output, err
		}

		output = u
	}

	return output, nil
}

func DeleteUser(id int) error {

	query := `
		DELETE FROM users WHERE user_id = $1
	`

	res, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	r, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if r == 0 {
		return errors.New("no selected id exist")
	}

	return nil
}
