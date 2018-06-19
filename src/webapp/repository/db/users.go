package db

import (
	"fmt"
)

// Users struct model from DB
type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//GetUsersDB get all users from DB
func GetUsersDB() []Users {
	var users []Users
	dbc := NewDbconnection()

	rows, err := dbc.Query("SELECT * FROM users")
	CheckErr(err)

	for rows.Next() {
		var user Users
		err = rows.Scan(&user.ID, &user.Name)
		CheckErr(err)
		fmt.Println("id:", user.ID, "Name:", user.Name)
		users = append(users, user)
	}
	return users
}
