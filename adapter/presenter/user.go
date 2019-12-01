package presenter

import (
	"net/http"

	"github.com/koheyeng/go_clean_architecture_sample/domain"
	"github.com/koheyeng/go_clean_architecture_sample/usecase/user"
	"golang.org/x/xerrors"
)

type userOutputPort struct {
	respWriter http.ResponseWriter
}

func NewUserOutputPort(w http.ResponseWriter) user.UserOutputPort {
	return &userOutputPort{
		respWriter: w,
	}
}

func (u *userOutputPort) RespUsers(users []*domain.User) error {
	if err := JSON(u.respWriter, http.StatusOK, users); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil
}
