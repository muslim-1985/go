package models

import (
	"database/sql"
	"errors"
	//"errors"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//var tableName = "users"

func (p *User) GetUser(db *sql.DB) error {
	return db.QueryRow("SELECT username, email FROM users WHERE id=$1",
		p.ID).Scan(&p.Username, &p.Email)
}

func (p *User) UpdateUser(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE users SET username=$1, email=$3 WHERE id=$3",
			p.Username, p.Email, p.ID)
	return err
}

func (p *User) DeleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", p.ID)
	return err
}

func (p *User) UserRegister(db *sql.DB) error {
	var checkUserExist *bool
	err := db.QueryRow("select exists(select email from users where email=$1)",
		p.Email).Scan(&checkUserExist)
	if *checkUserExist {
		return errors.New("A user is already registered to this mail")
	}
	if err != nil {
		return err
	}
	error1 := db.QueryRow(
		"INSERT INTO users(username, email, password) VALUES($1, $2, $3) RETURNING id", p.Username,
		p.Email, p.Password).Scan(&p.ID)
	if error1 != nil {
		return error1
	}
	return nil
}

func GetUsers(db *sql.DB, start, count int) ([]User, error) {
	rows, err := db.Query(
		"SELECT id, username FROM users LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var p User
		if err := rows.Scan(&p.ID, &p.Username); err != nil {
			return nil, err
		}
		users = append(users, p)
	}

	return users, nil
}
