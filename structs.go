package main

// Devices contains all LineageOS builds from the API.
type Devices map[string][]Build

// Build is a single build from LineageOS.
type Build struct {
	Date     string `json:"date"`
	DateTime int64  `json:"datetime"`
	FileName string `json:"filename"`
	FilePath string `json:"filepath"`
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
	Size     uint   `json:"size"`
	Type     string `json:"type"`
	Version  string `json:"version"`
}
