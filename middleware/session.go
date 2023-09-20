package middleware

import (
	"net/http"
	"strings"

	"github.com/tiptophelmet/nomess-core/v5/config"
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-core/v5/session"
	"github.com/tiptophelmet/nomess-core/v5/util"
)

func WithSession(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	var sessionToken string

	if headerSessionToken := r.Header.Get("Authorization"); headerSessionToken != "" {
		sessionToken = strings.TrimPrefix(headerSessionToken, "Bearer ")
	} else if cookieSessionToken, err := r.Cookie("session_token"); err == nil {
		sessionToken = cookieSessionToken.String()
	}

	isValid, err := session.Get().ValidateSessionToken(sessionToken)

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	if isValid && err != nil {
		rotatedToken, _ := session.Get().TryRotateSessionToken(sessionToken)

		if rotatedToken == "" {
			return w, r
		}

		jwtExpTime := config.Get("session.jwt.expiration.time").Required().Int()

		cookie.Value = rotatedToken
		cookie.MaxAge = jwtExpTime

		// Token successfully rotated
		http.SetCookie(w, &cookie)

		return w, r
	}

	logger.Debug("Unauthorized request to [%s] %s was halted", r.Method, util.GetRoutePattern(r))
	
	http.SetCookie(w, &cookie)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)

	return nil, r
}
