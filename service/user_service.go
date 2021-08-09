package user

import (
	"database/sql"
	"fmt"
	"log"

	db "go_sample/driver"
	"go_sample/entity"
)

// Service procides user's behavior
type Service struct{}

// User is alias of entity.User struct
type User entity.User

// get all User
func (s Service) GetAll() ([]User, error) {
	var user User
	users := make([]User, 0)
	db := db.GetDB()

	rows, err := db.Queryx("SELECT id, first_name, last_name FROM users;")

	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, err
}

// get a User
func (s Service) GetByID(id string) (User, error) {
	var user User
	db := db.GetDB()

	row := db.QueryRowx("SELECT id, first_name, last_name FROM users WHERE id=?;", id)
	err := row.StructScan(&user)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal("no user with id %d\n", id)
		log.Fatal(err)
	}

	return user, err
}

// create User model
func (s Service) CreateUser(firstname, lastname string) (User, error) {
	var user User
	db := db.GetDB()

	stmt, err := db.Prepare("INSERT users (first_name, last_name) VALUES(?, ?);")
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}

	res, err := stmt.Exec(firstname, lastname)
	defer stmt.Close()
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
	}

	id, _ := res.LastInsertId()

	row := db.QueryRowx("SELECT id, first_name, last_name FROM users WHERE id=?;", id)
	err = row.StructScan(&user)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal("no user with id %d\n", id)
		log.Fatal(err)
	}

	return user, err
}

// update a User
func (s Service) UpdateByID(id, firstname, lastname string) (User, error) {
	var user User
	db := db.GetDB()

	row := db.QueryRowx("SELECT id, first_name, last_name FROM users WHERE id=?;", id)
	err := row.StructScan(&user)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal("no user with id %d\n", id)
		log.Fatal(err)
	}

	stmt, err := db.Prepare("UPDATE users SET first_name=?, last_name=? WHERE id=?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(firstname, lastname, id)
	defer stmt.Close()
	if err != nil {
		fmt.Print(err.Error())
	}

	if err != nil {
		log.Fatal("Cannot run insert statement", err)
	}

	row = db.QueryRowx("SELECT id, first_name, last_name FROM users WHERE id=?;", id)
	err = row.StructScan(&user)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return user, err
}

// delete a User
func (s Service) DeleteByID(id string) error {
	var user User
	db := db.GetDB()

	row := db.QueryRowx("SELECT id, first_name, last_name FROM users WHERE id=?;", id)
	err := row.StructScan(&user)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal("no user with id %d\n", id)
		log.Fatal(err)
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id=?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	defer stmt.Close()
	if err != nil {
		fmt.Print(err.Error())
	}

	return err
}
