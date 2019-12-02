package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/koheyeng/go_clean_architecture_sample/adapter/repository"
	"github.com/koheyeng/go_clean_architecture_sample/usecase"
)

type UserHandler struct {
	dbHandler repository.DBHandler
}

func NewUser(dbHandler repository.DBHandler) *UserHandler {
	return &UserHandler{
		dbHandler: dbHandler,
	}
}

func (u *UserHandler) GetUserAge(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
	}

	inputPort, err := initInputPort(u.dbHandler, w)
	if err != nil {
	}
	err = inputPort.GetUserAge(usecase.UsersDto{
		ID: id,
	})
	if err != nil {
		log.Printf("%v\n", err)
	}
}
