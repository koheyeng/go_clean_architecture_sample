package usecase

import (
	"fmt"

	"github.com/koheyeng/go_clean_architecture_sample/domain"
)

type UsersInputPort interface {
	GetUserAge(UsersDto) error
}

type UsersRepository interface {
	GetUser(UsersDto) (*domain.Users, error)
}
type OuterSystemGateway interface {
	Send(userID int) error
}

type UsersPresenter interface {
	RespUserAge(int) error
}

type User struct {
	repo      UsersRepository
	gateway   OuterSystemGateway
	presenter UsersPresenter
}

func NewUser(repo UsersRepository, gateway OuterSystemGateway, presenter UsersPresenter) UsersInputPort {
	return &User{
		repo:      repo,
		gateway:   gateway,
		presenter: presenter,
	}
}

type UsersDto struct {
	ID int
}

func (u *User) GetUserAge(dto UsersDto) error {
	user, err := u.repo.GetUser(dto)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	if err := u.gateway.Send(dto.ID); err != nil {
		return fmt.Errorf(": %w", err)
	}

	if err := u.presenter.RespUserAge(user.GetAge()); err != nil {
		return fmt.Errorf(": %w", err)
	}

	return nil
}
