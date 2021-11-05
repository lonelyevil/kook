package khl

import (
	"bytes"
	"net/http"
	"time"
)

// New creates a khl session with default settings
func New(token string, l Logger, o ...SessionOption) (s *Session) {
	s = &Session{
		Client:       &http.Client{Timeout: 30 * time.Second},
		sequence:     new(int64),
		MaxRetry:     3,
		RetryTimeout: 60 * time.Second,
		ContentType:  "application/json",
		Logger:       l,
		snStore:      newBloomSnStore(),
		Sync:         true,
	}
	s.Identify.Token = "Bot " + token
	s.Identify.Compress = true
	for _, item := range o {
		item(s)
	}
	return
}

// SessionOption is the optional arguments for creating a session.
type SessionOption func(*Session)

// SessionWithVerifyToken adds the token for verifying webhook request.
func SessionWithVerifyToken(token string) SessionOption {
	return func(session *Session) {
		session.Identify.VerifyToken = token
	}
}

// SessionWithEncryptKey adds the key for decrypting webhook request.
func SessionWithEncryptKey(key []byte) SessionOption {
	return func(session *Session) {
		if len(key) < 32 {
			buf := &bytes.Buffer{}
			buf.Grow(32)
			buf.Write(key)
			key = buf.Bytes()
		}
		session.Identify.WebsocketKey = key
	}
}
