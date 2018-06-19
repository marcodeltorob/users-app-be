package controller

import (
	"net/http"
	"webapp/service"
)

type mainController struct {
	srv *service.UsersService
}

func (mc *mainController) MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//fmt.Fprint(w, mc.srv.MainService())
	w.Write(mc.srv.UsersService())
	w.WriteHeader(http.StatusOK)

}

// NewMainController ...
func NewMain(srv *service.UsersService) *mainController {
	return &mainController{srv}
}
