package router

import (
	"webapp/controller"
	"webapp/service"

	"github.com/gorilla/mux"
)

//Initrouter initialize all routers from application
func Initrouter() *mux.Router {

	r := mux.NewRouter()

	// Services
	s := service.NewUsersService()

	//Controllers
	c := controller.NewMain(s)

	r.HandleFunc("/users/", c.MainHandler).Methods("GET")
	return r
}
