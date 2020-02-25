// builds.go

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    builds, err := UnmarshalBuilds(bytes)
//    bytes, err = builds.Marshal()

package lineageos

import "encoding/json"

// Builds is the raw output from the LineageOS API.
type Builds map[string][]Build

// UnmarshalBuilds unmarshal LineageOS builds
func UnmarshalBuilds(data []byte) (builds Builds, err error) {
	err = json.Unmarshal(data, &builds)
	return builds, err
}

// Marshal marshal LineageOS builds.
func (r *Builds) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
