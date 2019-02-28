package main

// Devices contains all LineageOS builds from the API.
type Devices map[string][]Build

// Build is a single build from LineageOS.
type Build struct {
	Date     string `json:"date"`
	Datetime int64  `json:"datetime"`
	Filename string `json:"filename"`
	Filepath string `json:"filepath"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
	Version  string `json:"version"`
}

// HashedBuild is a build and its hash.
type HashedBuild struct {
	Worker int
	Build  *Build
	Hash   string
}
