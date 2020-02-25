package lineageos

import (
	"net/http"
)

// Service is a LineageOS Service.
type Service struct {
	agent  string
	client *http.Client
}

// New returns a new LineageOS service.
func New(agent string, client *http.Client) *Service {
	return &Service{
		agent:  agent,
		client: client,
	}
}
