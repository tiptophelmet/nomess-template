package errs

import "errors"

// Business-related errors
var (
	ErrUserInsert               = errors.New("user could not be saved")
	ErrUserVerificationInsert   = errors.New("user verification could not be prepared")
	ErrVerificationEmailNotSent = errors.New("verification email not sent")
)
