package models

import "time"

type Users struct {
	ID         int       `gorm:"type:integer;primary_key"`
	FirstName  string    `gorm:"type:varchar(36);not null"`
	FamilyName string    `gorm:"type:varchar(36);not null"`
	Age        int       `gorm:"type:integer;not null"`
	BirthDay   time.Time `gorm:"type:timestamp"`
	Address    string    `gorm:"type:varchar(128)"`
}
