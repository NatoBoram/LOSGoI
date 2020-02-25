package lineageos

import (
	"io"
	"net/http"
)

// request creates a new requests with basic headers.
func (s *Service) request(method string, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	s.headers(req)

	return req, err
}
