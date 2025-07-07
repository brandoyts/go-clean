package domain

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `bson:"_id" json:"id,omitempty"`
	Email     string    `bson:"email" json:"email,omitempty"`
	Password  string    `bson:"password" json:"password,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

var (
	ErrUserHashPassword = fmt.Errorf("cannot hash empty password")
)

func (u *User) HashPassword() error {
	if strings.TrimSpace(u.Password) == "" {
		return ErrUserHashPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}
