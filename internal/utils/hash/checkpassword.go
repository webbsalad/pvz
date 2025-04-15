package hash

import (
	"github.com/webbsalad/pvz/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return model.ErrPasswordMismatch
	}
	return nil
}
