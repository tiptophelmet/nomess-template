package session

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/errs"
	"github.com/tiptophelmet/nomess/logger"
)

func IssueSessionToken(userID string) (string, error) {
	jwtExpTime, err := config.Int64("session.jwt.expiration.time")
	if err != nil {
		logger.Crit("could not resolve session.jwt.expiration.time")
		return "", errs.ErrJwtNotIssued
	}

	// TODO: Add jti support
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Unix() + jwtExpTime,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret, err := config.Str("session.jwt.secret")
	if err != nil {
		logger.Crit("could not resolve session.jwt.secret")
		return "", errs.ErrJwtNotIssued
	}

	signedJwt, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		logger.Crit(fmt.Sprintf("failed to sign jwt: %v", err.Error()))
		return "", errs.ErrJwtNotIssued
	}

	return signedJwt, nil
}

func checkRotationCondition(jwtExpTime int64) (bool, error) {
	jwtExpWindow, err := config.Int64("session.jwt.expiration.window")
	if err != nil {
		logger.Crit("could not resolve session.jwt.expiration.window")
		return false, errs.ErrJwtNotIssued
	}

	currentTime := time.Now().Unix()

	isExpired := currentTime > jwtExpTime
	isRotationWindowOk := currentTime < jwtExpTime+jwtExpWindow

	return isExpired && isRotationWindowOk, nil
}

func TryRotateSessionToken(signedToken string) (string, error) {
	jwtSecret, err := config.Str("session.jwt.secret")
	if err != nil {
		logger.Crit("could not resolve session.jwt.secret")
		return "", errs.ErrJwtNotIssued
	}

	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if !parsedToken.Valid {
		return "", errs.ErrInvalidJwt
	}

	_, isTokenExpError := err.(*jwt.TokenExpiredError)

	if err != nil && !isTokenExpError {
		return "", errs.ErrJwtNotParsed
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errs.ErrInvalidJwtClaims
	}

	jwtExpTime := claims["exp"].(int64)

	isRotationOk, err := checkRotationCondition(jwtExpTime)
	if err != nil {
		return "", err
	}

	if isRotationOk {
		return IssueSessionToken(claims["sub"].(string))
	}

	return "", nil
}

func ValidateSessionToken(token string) (bool, error) {
	jwtSecret, err := config.Str("session.jwt.secret")
	if err != nil {
		logger.Crit("could not resolve session.jwt.secret")
		return false, errs.ErrJwtNotParsed
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if !parsedToken.Valid || err != nil {
		return false, nil
	}

	return true, nil
}
