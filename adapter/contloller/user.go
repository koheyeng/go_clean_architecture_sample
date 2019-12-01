package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/koheyeng/go_clean_architecture_sample/usecase/user"
)

type User struct {
	repositoryFactory func(user.DBHandler) user.UserRepository
	outputFactory     func(http.ResponseWriter) user.UserOutputPort
	inputFactory      func(user.UserRepository, user.UserOutputPort) user.UserInputPort
	dbHandler         user.DBHandler
}

func NewUser(dbHandler user.DBHandler, apiBaseURL string) *User {
	return &User{
		inputFactory: user.NewUser,
		dbHandler:    dbHandler,
	}
}

func (u *User) GetUsers(w http.ResponseWriter, r *http.Request) {

	repository := u.repositoryFactory(u.dbHandler)
	outputPort := u.outputFactory(w)
	inputPort := u.inputFactory(repository, outputPort)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {

	}
	name := r.URL.Query().Get("name")

	err = inputPort.GetUsers(user.UserDto{
		ID:   id,
		Name: name,
	})
	if err != nil {
		log.Printf("%v\n", err)
	}
}
