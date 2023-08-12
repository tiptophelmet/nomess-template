package password

import (
	"fmt"

	"github.com/tiptophelmet/nomess/logger"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		logger.Crit(fmt.Sprintf("could not hash a password: %s", err.Error()))
		return "", err
	}
	return string(hash), nil
}
