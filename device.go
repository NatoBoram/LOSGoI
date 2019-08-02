package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
)

// Devices contains all LineageOS builds from the API.
type Devices map[string][]*Build

// Name every builds from every devices.
func (devices *Devices) Name() {
	for device, builds := range *devices {
		for _, build := range builds {
			build.Device = device
		}
	}
}

// Count every builds from every devices.
func (devices *Devices) Count() (count int) {
	for _, builds := range *devices {
		for range builds {
			count++
		}
	}
	return
}

// Trim trims devices that were already hashed and those that aren't the latest of the device.
func (devices *Devices) Trim() (builds []*Build) {

	// Select only the newest build by device
	var newestBuilds []*Build
	for _, device := range *devices {
		var latest time.Time

		// Get the max date for this device
		for _, build := range device {
			if latest.Before(build.Datetime.Time) {
				latest = build.Datetime.Time
			}
		}

		// Put the build corresponding to the max date in the array
		for _, build := range device {
			if latest.Equal(build.Datetime.Time) {
				newestBuilds = append(newestBuilds, build)
			} else {
				fmt.Println("Skipping", aurora.Green(build.Filename).String()+".")
			}
		}
	}

	// Compare with the database
	for _, build := range newestBuilds {

		// Check if it needs to be hashed.
		bh, err := build.Select()
		if err == sql.ErrNoRows {
			builds = append(builds, build)
		} else if err != nil {
			fmt.Println("Couldn't select build", aurora.Green(build.Filename).String()+".")
			fmt.Println(err.Error())
		} else {
			fmt.Println("Skipping", aurora.Green(bh.Build.Filename).String()+".")
		}
	}

	return
}

// Hash every builds from every devices.
func (devices *Devices) Hash() {

	// Skip known builds
	builds := devices.Trim()

	// Progression
	index := 1
	total := len(builds)
	start := time.Now()

	for _, build := range builds {
		bh, err := build.Hash(index, total)
		if err != nil {
		} else if bh != nil && bh.Build != nil {
			bh.Save()
		}

		// Estimated Time Left
		fmt.Println("Estimated Time Left :", aurora.Bold(eta(start, index, total)).String()+".")

		index++
	}

	// Duration
	fmt.Println("Hashed in", aurora.Bold(time.Since(start).String()).String()+".")
}
