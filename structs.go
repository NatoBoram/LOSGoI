package main

import (
	"time"
)

// Devices contains all LineageOS builds from the API.
type Devices map[string][]*Build

// BuildDate is the `Date` inside a `Build`.
type BuildDate time.Time

// BuildDateTime is the `Datetime` inside a `Build`.
type BuildDateTime time.Time

// Build is a single build from LineageOS.
type Build struct {
	Date     BuildDate     `json:"date"`
	Datetime BuildDateTime `json:"datetime"`
	Filename string        `json:"filename"`
	Filepath string        `json:"filepath"`
	Sha1     string        `json:"sha1"`
	Sha256   string        `json:"sha256"`
	Size     int           `json:"size"`
	Type     string        `json:"type"`
	Version  string        `json:"version"`
}

// BuildHash is a build and its hash.
type BuildHash struct {
	Build *Build
	IPFS  string
}

// Database hosts the database configuration.
type Database struct {
	User     string
	Password string
	Address  string
	Port     int
	Database string
}
