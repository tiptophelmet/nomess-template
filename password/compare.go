package password

import (
	"fmt"

	"github.com/tiptophelmet/nomess/logger"

	"golang.org/x/crypto/bcrypt"
)

func CompareToHash(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	} else if err != nil {
		logger.Crit(fmt.Sprintf("could not compare password to hash: %s", err.Error()))
		return false
	}

	return true
}
