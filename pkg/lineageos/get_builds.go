package lineageos

import (
	"io/ioutil"
)

// GetBuilds gets all LineageOS builds from the API.
func (s *Service) GetBuilds() (builds Builds, err error) {

	// Request
	req, err := s.request("GET", apiBuilds, nil)
	if err != nil {
		return nil, err
	}

	// Response
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Bytes
	resb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Builds
	builds, err = UnmarshalBuilds(resb)
	if err != nil {
		return builds, err
	}

	return builds, err
}
