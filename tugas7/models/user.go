package model

import "errors"

var Users = []*User{}

type User struct {
	ID       int    `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username;unique;not null"`
	Password string `json:"password" gorm:"column:password;unique;not null"`
	Email    string `json:"email" gorm:"column:email"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("invalid username input")
	}

	if u.Email == "" {
		return errors.New("invalid username input")
	}

	if u.Password == "" {
		return errors.New("invalid password input")
	}

	return nil
}
