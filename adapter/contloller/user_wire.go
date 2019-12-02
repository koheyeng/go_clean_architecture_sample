//+build wireinject

package controller

import (
	"net/http"

	"github.com/koheyeng/go_clean_architecture_sample/adapter/gateway"
	"github.com/koheyeng/go_clean_architecture_sample/adapter/presenter"
	"github.com/koheyeng/go_clean_architecture_sample/adapter/repository"
	"github.com/koheyeng/go_clean_architecture_sample/usecase"

	"github.com/google/wire"
)

func initInputPort(dbHandler repository.DBHandler, w http.ResponseWriter) (usecase.UsersInputPort, error) {
	wire.Build(
		usecase.NewUser,
		repository.NewUsersRepository,
		gateway.NewOuterSystemGateway,
		presenter.NewUserOutputPort,
	)

	return &usecase.User{}, nil
}
