package losgoi

import (
	"bytes"
	"encoding/gob"
)

// Device is multiple builds from the same LineageOS device.
type Device struct {
	Name   string
	Builds *[]Build
}

// MarshalGob encodes a device in gob.
func (device *Device) MarshalGob() ([]byte, error) {
	b := new(bytes.Buffer)
	err := gob.NewEncoder(b).Encode(device)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), err
}

// UnmarshalGob decodes a device from gob.
func (device *Device) UnmarshalGob(data []byte) error {
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(device)
}
