package repository

import (
	"github.com/golang/xerrors"
	"github.com/koheyeng/go_clean_architecture_sample/domain"
	"github.com/koheyeng/go_clean_architecture_sample/usecase/users"
)

type usersRepository struct {
	dbHandler DBHandler
}

type DBHandler interface {
	Query() (*domain.Users, error)
	ByID(int) DBHandler
	ByName(string, string) DBHandler
}

func NewUsersRepository(dbHandler DBHandler) users.UsersRepository {
	return &usersRepository{
		dbHandler: dbHandler,
	}
}

func (e *usersRepository) GetUser(dto users.UsersDto) (*domain.Users, error) {

	user, err := e.dbHandler.ByID(dto.ID).Query()
	if err != nil {
		return nil, xerrors.Errorf("edges list acquisition error: %w", err)
	}

	return user, nil
}
