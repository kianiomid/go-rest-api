package models

import (
	"projects/go-rest-api/src/api/security"
	"time"
)

type User struct {
	ID        int32    `gorm:"primary_key;auto_increment" json:"id" `
	Nickname  string    `gorm:"size:20;not null;unique" json:"nickname" `
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null;" json:"password" `
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func (u *User) BeforeSave() error {

	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}
