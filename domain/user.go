package domain

import "time"

type User struct {
	ID       int       `gorm:"type:integer;primary_key"`
	Name     string    `gorm:"type:varchar(36);not null"`
	Age      int       `gorm:"type:integer;not null"`
	BirthDay time.Time `gorm:"type:timestamp"`
	Address  string    `gorm:"type:varchar(128)"`
}
