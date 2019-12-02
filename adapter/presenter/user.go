package presenter

import (
	"net/http"

	"github.com/koheyeng/go_clean_architecture_sample/usecase/users"
	"golang.org/x/xerrors"
)

type usersPresenter struct {
	respWriter http.ResponseWriter
}

func NewUserOutputPort(w http.ResponseWriter) users.UsersPresenter {
	return &usersPresenter{
		respWriter: w,
	}
}

type UserAge struct {
	Age int
}

func (u *usersPresenter) RespUserAge(age int) error {
	if err := JSON(u.respWriter, http.StatusOK, UserAge{age}); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil
}
