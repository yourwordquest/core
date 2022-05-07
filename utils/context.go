package utils

import (
	"context"
	"time"
)

// TimedContext is useful when you want to limit an execution by time.
func TimedContext(seconds time.Duration) context.Context {
	ct, _ := context.WithTimeout(context.Background(), seconds*time.Second)
	return ct
}
