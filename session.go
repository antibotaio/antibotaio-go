package antibotaio

import (
	"net/http"
	"time"
)

type Session struct {
	APIKey string
	client *http.Client
}

func NewSession(apiKey string) *Session {
	return &Session{
		APIKey: apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *Session) WithClient(client *http.Client) *Session {
	s.client = client
	return s
}
