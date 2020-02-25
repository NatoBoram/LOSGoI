package lineageos

import (
	"strconv"
	"strings"
	"time"
)

// Build is a LineageOS build.
type Build struct {
	Date     BuildDate     `json:"date"`
	Datetime BuildDateTime `json:"datetime"`
	Filename string        `json:"filename"`
	Filepath string        `json:"filepath"`
	Recovery *Recovery     `json:"recovery,omitempty"`
	Sha1     string        `json:"sha1"`
	Sha256   string        `json:"sha256"`
	Size     int64         `json:"size"`
	Type     Type          `json:"type"`
	Version  string        `json:"version"`
}

// BuildDate is the `Date` inside a `Build`.
type BuildDate time.Time

// UnmarshalJSON parses a formatted string and returns the time value it represents.
func (bd *BuildDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*bd = BuildDate(t)
	return err
}

// BuildDateTime is the `Datetime` inside a `Build`.
type BuildDateTime time.Time

// UnmarshalJSON parses a formatted string and returns the time value it represents.
func (bdt *BuildDateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	sec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	epochTime := time.Unix(sec, 0).In(time.UTC)
	*bdt = BuildDateTime(epochTime)
	return nil
}
