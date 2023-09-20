package util

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-template/app/logformat"
)

type RequestIdContextKey string

const key RequestIdContextKey = "NOMESS-REQUEST-ID"

func IssueRequestID(r *http.Request) *http.Request {
	rid := uuid.New().String()

	if logger.Logger() != nil {
		ridFormat := &logformat.RequestIdFormatter{}
		ridFormat.SetRequestId(rid)

		logger.Logger().SetFormatter(ridFormat)
	}

	reqIdCtx := context.WithValue(r.Context(), key, rid)
	return r.WithContext(reqIdCtx)
}

func RequestID(r *http.Request) (string, error) {
	reqIDContextVal, toStrOk := r.Context().Value(key).(string)
	if !toStrOk {
		err := fmt.Errorf("no request ID was issued for %s", r.URL.RequestURI())

		logger.Error(err.Error())
		return "", err
	}

	return reqIDContextVal, nil
}
