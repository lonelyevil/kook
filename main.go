package khl

import (
	"net/http"
	"time"
)

// New creates a khl session with default settings
func New(token string, l Logger) (s *Session) {
	s = &Session{
		Client:       &http.Client{Timeout: 30 * time.Second},
		sequence:     new(int64),
		MaxRetry:     3,
		RetryTimeout: 60 * time.Second,
		ContentType:  "application/json",
		Logger:       l,
	}
	s.Identify.Token = "Bot " + token
	s.Identify.Compress = true
	return
}
