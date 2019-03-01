package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

// IsEmpty returns `true` if one field from a `Database` object is empty.
func (database Database) IsEmpty() bool {
	return database.Address == "" ||
		database.Database == "" ||
		database.Password == "" ||
		database.Port == 0 ||
		database.User == ""
}

// ToConnectionString creates a connection string from a `Database`.
func (database Database) ToConnectionString() string {
	return database.User + ":" + database.Password + "@tcp(" + database.Address + ":" + strconv.Itoa(database.Port) + ")/" + database.Database + "?parseTime=true"
}

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
func (bd BuildDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(bd)
}

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
func (bdt BuildDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(bdt)
}

// Name every builds from every devices.
func (devices Devices) Name() {
	for device, builds := range devices {
		for _, build := range builds {
			build.Device = device
		}
	}
}

// Count every builds from every devices.
func (devices Devices) Count() (count int) {
	for _, builds := range devices {
		for range builds {
			count++
		}
	}
	return
}

// Hash every builds from every devices.
func (devices Devices) Hash() {
	index := float64(1)
	total := float64(devices.Count())

	for _, builds := range devices {
		for _, build := range builds {

			// Check if it needs to be hashed.
			bh, err := build.Select()
			if err == sql.ErrNoRows {
				build.Hash(index, total)
			} else if err != nil {
				fmt.Println("Couldn't select build", aurora.Green(build.Filename).String()+".")
				fmt.Println(err.Error())
			} else {
				fmt.Println("Skipping", aurora.Green(bh.Build.Filename).String()+".")
			}

			index++
		}
	}
}

// Select this build from the database.
func (build Build) Select() (bh BuildHash, err error) {

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
func (build Build) Hash(index float64, total float64) {

	// Log
	fmt.Println("Processing the build", aurora.Green(build.Filename).String()+".")
	start := time.Now()

	// Add URL to ipfs without storing the data locally.
	filepath := mirrorbits + build.Filepath
	out, err := exec.Command("ipfs", "urlstore", "add", filepath).Output()
	if err != nil {
		fmt.Println("Failed to execute the command.")
		fmt.Println(aurora.Bold("Command :"), "ipfs", "urlstore", "add", aurora.Blue(filepath))
		fmt.Println(err.Error())
		return
	}

	// Log
	fmt.Println(aurora.Bold(fmt.Sprintf("%3.2f%%", index/total*100)), "|", aurora.Green(build.Filename), "|", aurora.Cyan(string(out)), "|", time.Since(start).String())

	// Finished
	go BuildHash{
		Build: &build,
		IPFS:  string(out),
	}.Save()
}

// Save the BuildHash to the database.
func (buildHash BuildHash) Save() {
	_, err := db.Exec("insert into `builds`(`device`, `date`, `datetime`, `filename`, `filepath`, `sha1`, `sha256`, `size`, `type`, `version`, `ipfs`) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", buildHash.Build.Device, time.Time(buildHash.Build.Date), time.Time(buildHash.Build.Datetime), buildHash.Build.Filename, buildHash.Build.Filepath, buildHash.Build.Sha1, buildHash.Build.Sha256, buildHash.Build.Size, buildHash.Build.Type, buildHash.Build.Version, buildHash.IPFS)
	if err != nil {
		fmt.Println("Couldn't save build", aurora.Green(buildHash.Build.Filename).String()+".")
		fmt.Println(err.Error())
		return
	}
}

// Pin a build to the local IPFS gateway.
func (buildHash BuildHash) Pin() {
	out, err := exec.Command("ipfs-cluster-ctl", "pin", "add", buildHash.IPFS, "--name", buildHash.Build.Filename).CombinedOutput()
	if err != nil {
		fmt.Println("Failed to pin a build.")
		fmt.Println(aurora.Bold("Command :"), "ipfs-cluster-ctl", "pin", "add", aurora.Cyan(buildHash.IPFS), "--name", aurora.Green(buildHash.Build.Filename))
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}

// Unpin a build to the local IPFS gateway.
func (buildHash BuildHash) Unpin() {
	out, err := exec.Command("ipfs-cluster-ctl", "pin", "rm", buildHash.IPFS).CombinedOutput()
	if err != nil {
		fmt.Println("Failed to pin a build.")
		fmt.Println(aurora.Bold("Command :"), "ipfs-cluster-ctl", "pin", "rm", aurora.Cyan(buildHash.IPFS))
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}
