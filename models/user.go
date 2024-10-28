package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	Parent UserRole = "parent"
	Kid UserRole = "kid"
	Player UserRole = "player"
)

type User struct {
	ID uint `json:"id" gorm:"primarykey"`
	Email string `json:"email" gorm:"unique"`
	Role UserRole `json:"role" gorm:"text;default:player"`
	Password string `json:"-"` // Do not compute the password in json
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"udpated_at"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID == 1 {
		db.Model(u).Update("role", Parent)
	}
	return
}