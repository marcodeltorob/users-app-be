package service

import (
	"encoding/json"
	"log"
	"webapp/repository/db"
)

type UsersService struct {
}

//NewUsersService ...
func NewUsersService() *UsersService {
	return &UsersService{}
}

// UsersService returns Users JSON
func (ms *UsersService) UsersService() []byte {
	users := db.GetUsersDB()

	jsonObj, err := json.Marshal(users)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	//return string(jsonObj)
	return jsonObj

}
