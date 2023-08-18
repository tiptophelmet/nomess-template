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
	jwtExpTime := config.Get("session.jwt.expiration.time").Required().Int64()

	// TODO: Add jti support
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Unix() + jwtExpTime,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := config.Get("session.jwt.secret").Required().Str()

	signedJwt, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		logger.Crit(fmt.Sprintf("failed to sign jwt: %v", err.Error()))
		return "", errs.ErrJwtNotIssued
	}

	return signedJwt, nil
}

func checkRotationCondition(jwtExpTime int64) (bool, error) {
	jwtExpWindow := config.Get("session.jwt.expiration.window").Required().Int64()
	currentTime := time.Now().Unix()

	isExpired := currentTime > jwtExpTime
	isRotationWindowOk := currentTime < jwtExpTime+jwtExpWindow

	return isExpired && isRotationWindowOk, nil
}

func TryRotateSessionToken(signedToken string) (string, error) {
	jwtSecret := config.Get("session.jwt.secret").Required().Str()

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
	jwtSecret := config.Get("session.jwt.secret").Required().Str()

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if !parsedToken.Valid || err != nil {
		return false, nil
	}

	return true, nil
}
