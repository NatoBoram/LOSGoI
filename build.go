package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

// Build is a single build from LineageOS.
type Build struct {
	Device   string
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
	return nil
}

// MarshalJSON returns the JSON encoding of a `BuildDateTime`.
func (bd *BuildDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(bd)
}

// BuildDateTime is the `Datetime` inside a `Build`.
type BuildDateTime time.Time

// UnmarshalJSON parses a formatted string and returns the time value it represents.
func (bdt *BuildDateTime) UnmarshalJSON(b []byte) error {
	t := strings.Trim(string(b), `"`)
	sec, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return err
	}
	epochTime := time.Unix(sec, 0)
	*bdt = BuildDateTime(epochTime)
	return nil
}

// MarshalJSON returns the JSON encoding of a `BuildDateTime`.
func (bdt *BuildDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(bdt)
}

// Select this build from the database.
func (build *Build) Select() (bh BuildHash, err error) {

	// Prevent null pointers
	bh = BuildHash{Build: &Build{}}
	var t1 time.Time
	var t2 time.Time

	err = db.QueryRow("select `device`, `date`, `datetime`, `filename`, `filepath`, `sha1`, `sha256`, `size`, `type`, `version`, `ipfs` from `builds` where `device` = ? and `date` = ? and `datetime` = ? and `filename` = ? and `filepath` = ? and `sha1` = ? and `sha256` = ? and `size` = ? and `type` = ? and `version` = ?;", build.Device, time.Time(build.Date), time.Time(build.Datetime), build.Filename, build.Filepath, build.Sha1, build.Sha256, build.Size, build.Type, build.Version).Scan(&bh.Build.Device, &t1, &t2, &bh.Build.Filename, &bh.Build.Filepath, &bh.Build.Sha1, &bh.Build.Sha256, &bh.Build.Size, &bh.Build.Type, &bh.Build.Version, &bh.IPFS)

	bh.Build.Date = BuildDate(t1)
	bh.Build.Datetime = BuildDateTime(t2)

	return
}

// Hash this build then save it.
func (build *Build) Hash(index int, total int) (bh *BuildHash, err error) {

	// Log
	fmt.Println("Processing build", aurora.Green(build.Filename).String()+"...")
	start := time.Now()

	// Add the file to IPFS
	filepath := mirrorbits + build.Filepath
	out, err := exec.Command("ipfs-cluster-ctl", "add", "-w", "--chunker=rabin", "--name", build.Filename, "--replication-min", build.RMin(), "--replication-max", build.RMax(), filepath).Output()
	if err != nil {
		fmt.Println("Failed to download a build.")
		fmt.Println(aurora.Bold("Command :"), "ipfs-cluster-ctl", "add", "-w", "--chunker=rabin", "--name", aurora.Green(build.Filename), "--replication-min", build.RMin(), "--replication-max", build.RMax(), aurora.Blue(filepath))

		// Log the error from the command
		ee, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println(string(ee.Stderr))
		}

		fmt.Println(string(out))
		return
	}

	// Remove garbage
	slice := strings.Fields(string(out))
	hash := slice[len(slice)-1]
	hash = strings.Trim(hash, "\n")

	if hash == hashEmptyFolder {

		// Check for "empty folder" hash
		unpin()
		return nil, errHashEmptyFolder
	} else if !strings.Contains(hash, "Qm") {

		// Check for invalid hash
		fmt.Println("Invalid hash :", hash)
		return nil, errInvalidHash
	}

	// Create the object
	bh = &BuildHash{
		Build: build,
		IPFS:  hash,
	}

	// Percentage
	fmt.Println(aurora.Bold(percent(index, total)), "|", aurora.Green(build.Filename), "|", aurora.Cyan(bh.IPFS), "|", time.Since(start).String())

	return
}

// RMin is the minimum replication factor for a given build.
func (build *Build) RMin() string {
	return strconv.Itoa(1)
}

// RMax is the maximum replication factor for a given build.
func (build *Build) RMax() string {
	return strconv.Itoa(build.Size/(speed*seconds) + 1)
}
