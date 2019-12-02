package domain

import (
	"errors"
	"fmt"
	"time"
)

type Users struct {
	id         int
	firstName  string
	familyName string
	birthDay   time.Time
	address    string
}

func NewUser(id int, firstName string, familyName string, birthDay time.Time, address string) (*Users, error) {
	if birthDay.After(time.Now()) {
		return nil, errors.New("Incorrect birthday Error")
	}

	return &Users{
		id:         id,
		firstName:  firstName,
		familyName: familyName,
		birthDay:   birthDay,
		address:    address,
	}, nil
}

func (u *Users) GetFullName() string {
	return fmt.Sprintf("%s %s", u.firstName, u.familyName)
}

func (u *Users) GetAge() int {
	return time.Now().Year() - u.birthDay.Year()
}
