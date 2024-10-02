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
	Email string `json:"email" gorm:"primarykey"`
	Role UserRole `json:"role" gorm:"text;default:player"`
	Password string `json:"-"` // Do not compute the password in json
	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time	`json:"udpatedAt"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID == 1 {
		db.Model(u).Update("role", Parent)
	}
	return
}