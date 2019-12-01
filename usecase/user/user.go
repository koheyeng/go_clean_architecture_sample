package user

import (
	"github.com/koheyeng/go_clean_architecture_sample/domain"
	"golang.org/x/xerrors"
)

type UserInputPort interface {
	GetUsers(UserDto) error
}

type UserRepository interface {
	GetUsers(UserDto) ([]*domain.User, error)
}

type UserOutputPort interface {
	RespUsers([]*domain.User) error
}

type DBHandler interface {
	Query() ([]*domain.User, error)
	ByID(int) DBHandler
	ByName(string) DBHandler
}

type User struct {
	repo       UserRepository
	outputPort UserOutputPort
}

func NewUser(repo UserRepository, outputPort UserOutputPort) UserInputPort {
	return &User{
		repo:       repo,
		outputPort: outputPort,
	}
}

type UserDto struct {
	ID   int
	Name string
}

func (u *User) GetUsers(dto UserDto) error {
	ee, err := u.repo.GetUsers(dto)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.outputPort.RespUsers(ee); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
