package lineageos

import (
	"net/http"
)

// headers sets the default headers for any given request.
func (s *Service) headers(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")
	req.Header.Set("User-Agent", s.agent)
}
