package middleware

import (
	"net/http"
	"strings"

	"github.com/tiptophelmet/nomess/internal/config"
	"github.com/tiptophelmet/nomess/internal/session"
)

func WithSession(w http.ResponseWriter, r *http.Request) {
	var sessionToken string

	if headerSessionToken := r.Header.Get("Authorization"); headerSessionToken != "" {
		sessionToken = strings.TrimPrefix(headerSessionToken, "Bearer ")
	} else if cookieSessionToken, err := r.Cookie("session_token"); err == nil {
		sessionToken = cookieSessionToken.String()
	}

	isValid, err := session.ValidateSessionToken(sessionToken)

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	if isValid && err != nil {
		rotatedToken, _ := session.TryRotateSessionToken(sessionToken)

		if rotatedToken == "" {
			return
		}

		jwtExpTime := config.Get("session.jwt.expiration.time").Required().Int()

		cookie.Value = rotatedToken
		cookie.MaxAge = jwtExpTime

		// Token successfully rotated
		http.SetCookie(w, &cookie)

		return
	}

	// Token is invalid - remove cookie, respond with unauthorized status
	http.SetCookie(w, &cookie)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
