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
	ErrPasswordHash                    = errors.New("password could not be hashed")
	ErrTypeAssertion                   = errors.New("type assertion from interface{} failed")
	ErrJwtNotIssued                    = errors.New("failed to issue jwt")
	ErrJwtNotParsed                    = errors.New("failed to parse jwt")
	ErrInvalidJwtClaims                = errors.New("jwt has invalid claims")
	ErrInvalidJwt                      = errors.New("jwt is invalid")
	ErrDbConnectionFailed              = errors.New("failed to connect to database")
	ErrCacheStoreConnectionFailed      = errors.New("failed to connect to cache store")
	ErrPubSubBrokerConnectionFailed    = errors.New("failed to connect to pubsub broker")
	ErrPubSubBrokerConnectionNotClosed = errors.New("failed to close connection to pubsub broker")
)
