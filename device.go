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

// Trim trims devices that were already hashed
func (devices Devices) Trim() (builds []*Build) {
	for _, device := range devices {
		for _, build := range device {

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
	}

	return
}

// Hash every builds from every devices.
func (devices Devices) Hash() {

	// Skip known builds
	builds := devices.Trim()

	// Progression
	index := 1
	total := len(builds)
	start := time.Now()

	for _, build := range builds {
		build.Hash(float64(index), float64(total))

		// Estimated Time Left
		fmt.Println("Estimated Time Left :", aurora.Bold(eta(start, index, total)).String()+".")

		index++
	}

	// Duration
	fmt.Println("Hashed in", aurora.Bold(time.Since(start).String()).String()+".")
}
