package logformat

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type RequestIdFormatter struct {
	requestId string
}

func (f *RequestIdFormatter) SetRequestId(requestId string) {
	f.requestId = strings.TrimSpace(requestId)
}

func (f *RequestIdFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var fmtLogStr string

	entryLevel := strings.ToUpper(entry.Level.String())
	entryTime := entry.Time.UTC().Format(time.DateTime)

	if f.requestId != "" {
		fmtLogStr = fmt.Sprintf("[%s][%s][%s] %s\n", f.requestId, entryLevel, entryTime, entry.Message)
	} else {
		fmtLogStr = fmt.Sprintf("[%s][%s] %s\n", entryLevel, entry.Time, entry.Message)
	}

	return []byte(fmtLogStr), nil

}
