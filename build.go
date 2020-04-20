package losgoi

import (
	"bytes"
	"encoding/gob"

	"gitlab.com/NatoBoram/LOSGoI/pkg/lineageos"
)

// Build is a single build from LineageOS.
type Build struct {
	Build  *lineageos.Build
	Device string
	IPFS   *string
}

// MarshalGob encodes a build in gob.
func (build *Build) MarshalGob() ([]byte, error) {
	b := new(bytes.Buffer)
	err := gob.NewEncoder(b).Encode(build)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), err
}

// UnmarshalGob decodes a build from gob.
func (build *Build) UnmarshalGob(data []byte) error {
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(build)
}
