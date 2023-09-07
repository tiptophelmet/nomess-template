package app

import (
	"github.com/tiptophelmet/nomess/internal/config"
	"github.com/tiptophelmet/nomess/internal/session"
)

func initSession() {
	jwtExpTime := config.Get("session.jwt.expiration.time").Required().Int64()
	jwtExpWindow := config.Get("session.jwt.expiration.window").Required().Int64()
	jwtSecret := config.Get("session.jwt.secret").Required().Str()

	session.Init(jwtExpTime, jwtExpWindow, jwtSecret)
}
