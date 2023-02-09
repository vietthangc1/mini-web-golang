package modules

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/vietthangc1/mini-web-golang/models"
)

func QueryAddUser(db *sql.DB, q string, u models.User) (models.User, error) {
	var newUser models.User

	id := time.Now().UnixMilli()
	newUser.ID = strconv.Itoa(int(id))

	stmt, err := db.Prepare(q)
	if err != nil {
		return models.User{}, err
	}

	res, err := stmt.Exec(
		u.ID,
		u.Email,
		u.Password,
	)
	if err != nil {
		return models.User{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return models.User{}, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return models.User{}, err
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	return u, nil

}

func QueryDeleteUser(db *sql.DB, q, id string) (error) {
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(
		id,
	)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return nil
}

func QueryGetUserByEmail (db *sql.DB, q string, email string) (models.User, error) {
	var userQuery models.User 

	err := db.QueryRow(q, email).Scan(
		&userQuery.ID,
		&userQuery.Email,
		&userQuery.Password,
	)

	if err != nil {
		return models.User{}, err
	}
	return userQuery, nil
}