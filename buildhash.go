package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"github.com/logrusorgru/aurora"
)

// BuildHash is a build and its hash.
type BuildHash struct {
	Build *Build
	IPFS  string
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

// Pin a build to the local IPFS cluster.
func (buildHash BuildHash) Pin() {
	const speed = 10 * 1024 * 1024
	const seconds = 60
	size := buildHash.Build.Size

	min := strconv.Itoa(1)
	max := strconv.Itoa(size/(speed*seconds) + 1)

	out, err := exec.Command("ipfs-cluster-ctl", "pin", "add", buildHash.IPFS, "--name", buildHash.Build.Filename, "--replication-min", min, "--replication-max", max).Output()
	if err != nil {
		fmt.Println("Failed to pin a build.")
		fmt.Println(aurora.Bold("Command :"), "ipfs-cluster-ctl", "pin", "add", aurora.Cyan(buildHash.IPFS), "--name", aurora.Green(buildHash.Build.Filename), "--replication-min", min, "--replication-max", max)

		// Log the error from the command
		ee, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println(string(ee.Stderr))
		}
	}

	fmt.Println(string(out))
}

// Unpin a build from the local IPFS cluster.
func (buildHash BuildHash) Unpin() {
	out, err := exec.Command("ipfs-cluster-ctl", "pin", "rm", buildHash.IPFS).Output()
	if err != nil {
		fmt.Println("Failed to unpin a build from the cluster.")
		fmt.Println(aurora.Bold("Command :"), "ipfs-cluster-ctl", "pin", "rm", aurora.Cyan(buildHash.IPFS))

		// Log the error from the command
		ee, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println(string(ee.Stderr))
		}
	}
	fmt.Println(string(out))
}
