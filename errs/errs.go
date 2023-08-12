package errs

import "errors"

// Business-related errors
var (
	ErrUserInsert               = errors.New("user could not be saved")
	ErrUserVerificationInsert   = errors.New("user verification could not be prepared")
	ErrVerificationEmailNotSent = errors.New("verification email not sent")
)

// App errors
var (
	ErrPasswordHash   = errors.New("password could not be hashed")
	ErrConfigNotFound = errors.New("config could not be found")
	ErrTypeAssertion  = errors.New("type assertion from interface{} failed")
)
