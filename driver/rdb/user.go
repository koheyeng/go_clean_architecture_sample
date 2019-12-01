package rdb

import (
	"github.com/jinzhu/gorm"
	"github.com/koheyeng/go_clean_architecture_sample/domain"
	"github.com/koheyeng/go_clean_architecture_sample/usecase/user"
	"golang.org/x/xerrors"
)

type scopeFunc struct {
	funcs []func(db *gorm.DB) *gorm.DB
}

type UserHandler struct {
	*gorm.DB
	scopeFunc
}

func (h *UserHandler) Query() ([]*domain.User, error) {
	res := []*domain.User{}
	if err := h.Scopes(h.funcs...).Order("id").Find(&res).Error; err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return res, nil
}

func (h *UserHandler) ByID(id int) user.DBHandler {
	h.funcs = append(h.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Where(domain.User{ID: id})
	})

	return h
}

func (h *UserHandler) ByName(name string) user.DBHandler {
	h.funcs = append(h.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Where(domain.User{Name: name})
	})

	return h
}
