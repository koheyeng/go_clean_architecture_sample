package gateway

import (
	"github.com/koheyeng/go_clean_architecture_sample/domain"
	"github.com/koheyeng/go_clean_architecture_sample/usecase/user"
	"golang.org/x/xerrors"
)

type userRepository struct {
	dbHandler user.DBHandler
}

func NewUserRepository(dbHandler user.DBHandler) user.UserRepository {
	return &userRepository{
		dbHandler: dbHandler,
	}
}

func (u *userRepository) GetUsers(dto user.UserDto) ([]*domain.User, error) {
	users, err := u.dbHandler.ByID(dto.ID).ByName(dto.Name).Query()
	if err != nil {
		return nil, xerrors.Errorf("user acquisition error: %w", err)
	}

	return users, nil
}
