package modules

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/vietthangc1/mini-web-golang/models"
)

func QueryAddUser(q string, u models.User) (models.User, error) {
	db, _ := sql.Open("mysql", os.Getenv("mysqlLogin"))
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

func QueryDeleteUser(q, id string) (error) {
	db, _ := sql.Open("mysql", os.Getenv("mysqlLogin"))
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

func QueryGetUserByEmail (q string, email string) (models.User, error) {
	db, _ := sql.Open("mysql", os.Getenv("mysqlLogin"))
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