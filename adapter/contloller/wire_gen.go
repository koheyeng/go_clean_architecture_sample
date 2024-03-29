// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controller

import (
	"github.com/koheyeng/go_clean_architecture_sample/adapter/gateway"
	"github.com/koheyeng/go_clean_architecture_sample/adapter/presenter"
	"github.com/koheyeng/go_clean_architecture_sample/adapter/repository"
	"github.com/koheyeng/go_clean_architecture_sample/usecase"
	"net/http"
)

// Injectors from user_wire.go:

func initInputPort(dbHandler repository.DBHandler, w http.ResponseWriter) (usecase.UsersInputPort, error) {
	usersRepository := repository.NewUsersRepository(dbHandler)
	outerSystemGateway := gateway.NewOuterSystemGateway()
	usersPresenter := presenter.NewUserOutputPort(w)
	usersInputPort := usecase.NewUsers(usersRepository, outerSystemGateway, usersPresenter)
	return usersInputPort, nil
}
