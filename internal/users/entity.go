package users

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	passwordMinLen = 6
	passwordMaxLen = 72
)

var (
	ErrNameRequired     = errors.New("Name is required")
	ErrLoginRequired    = errors.New("Login is required")
	ErrPasswordRequired = errors.New("Password is required")
	ErrPasswordTooShort = errors.New("Password must have at least 6 characters")
	ErrPasswordTooLong  = errors.New("Password must have less than 72 characters")
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"last_login"`
}

func (u *User) SetEncryptedPassword(password string) error {

	if password == "" {
		return ErrPasswordRequired
	}
	if len(password) < passwordMinLen {
		return ErrPasswordTooShort
	}
	if len(password) > passwordMaxLen {
		return ErrPasswordTooLong
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Error hashing the password")
	}

	u.Password = fmt.Sprintf("%x", encryptedPassword)

	return nil

}

func (u *User) Validate() error {

	if u.Name == "" {
		return ErrNameRequired
	}
	if u.Login == "" {
		return ErrLoginRequired
	}

	return nil
}

func New(name, login, password string) (*User, error) {

	u := User{Name: name, Login: login, Password: password, ModifiedAt: time.Now()}

	err := u.SetEncryptedPassword(password)
	if err != nil {
		return nil, err
	}

	err = u.Validate()
	if err != nil {
		return nil, err
	}

	return &u, nil
}
