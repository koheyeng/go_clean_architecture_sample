package rdb

import (
	"github.com/jinzhu/gorm"
	"github.com/koheyeng/go_clean_architecture_sample/adapter/repository"
	"github.com/koheyeng/go_clean_architecture_sample/db/models"
	"github.com/koheyeng/go_clean_architecture_sample/domain"
	"golang.org/x/xerrors"
)

type scopeFunc struct {
	funcs []func(db *gorm.DB) *gorm.DB
}

type UsersHandler struct {
	*gorm.DB
	scopeFunc
}

func NewUserHandler(db *gorm.DB) *UsersHandler {
	var sf scopeFunc
	return &UsersHandler{db, sf}
}

func (h *UsersHandler) Query() (*domain.Users, error) {
	m := &models.Users{}
	if err := h.Scopes(h.funcs...).Order("id").Find(m).Error; err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res, err := domain.NewUser(m.ID, m.FirstName, m.FamilyName, m.BirthDay, m.Address)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (h *UsersHandler) ByID(id int) repository.DBHandler {
	h.funcs = append(h.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Where(models.Users{ID: id})
	})

	return h
}

func (h *UsersHandler) ByName(firstName, familyName string) repository.DBHandler {
	h.funcs = append(h.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Where(models.Users{FirstName: firstName, FamilyName: familyName})
	})

	return h
}
