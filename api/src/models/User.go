package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare() error {
	err := user.validate()
	if err != nil {
		return err
	}

	user.format()

	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("Name é required and cannot be blank")
	}

	if user.Username == "" {
		return errors.New("Username é required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("Email é required and cannot be blank")
	}

	if user.Password == "" {
		return errors.New("Password é required and cannot be blank")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
}
