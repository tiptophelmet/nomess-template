package router

import (
	"fmt"

	"github.com/tiptophelmet/nomess/logger"
)

var lockedPatterns []string

func LockRoute(pattern string) {
	lockedPatterns = append(lockedPatterns, pattern)
}

func PeekRouteLock(assumed string) {
	for _, locked := range lockedPatterns {
		if matchLockedRoute(locked, assumed) {
			logger.Emergency(fmt.Sprintf("route reserved for core features & cannot be modified: %v", assumed))
		}
	}
}

func matchLockedRoute(locked string, assumed string) bool {
	return locked == assumed
}
