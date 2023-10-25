package util

import (
	"errors"
)

func PasswordValidator(password string) error {
	if len(password) < 8 {
		return errors.New("password minimal 8 character")
	}
	return nil
}
