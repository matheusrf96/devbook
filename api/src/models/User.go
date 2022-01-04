package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(step string) error {
	err := user.validate(step)
	if err != nil {
		return err
	}

	err = user.format(step)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name é required and cannot be blank")
	}

	if user.Username == "" {
		return errors.New("Username é required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("Email é required and cannot be blank")
	}

	err := checkmail.ValidateFormat(user.Email)
	if err != nil {
		return errors.New("Email is not valid")
	}

	if step == "register" && user.Password == "" {
		return errors.New("Password é required and cannot be blank")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}
