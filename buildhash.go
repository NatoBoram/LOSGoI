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
func (bh BuildHash) Save() {

	// Insert
	_, err := db.Exec("insert into `builds`(`device`, `date`, `datetime`, `filename`, `filepath`, `sha1`, `sha256`, `size`, `type`, `version`, `ipfs`) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", bh.Build.Device, time.Time(bh.Build.Date), time.Time(bh.Build.Datetime), bh.Build.Filename, bh.Build.Filepath, bh.Build.Sha1, bh.Build.Sha256, bh.Build.Size, bh.Build.Type, bh.Build.Version, bh.IPFS)
	if err != nil {
		fmt.Println("Couldn't save build", aurora.Green(bh.Build.Filename).String()+".")
		fmt.Println(err.Error())
		return
	}

	// Pin
	news, err := getLatestBuildsFromDevice(bh)
	if err != nil {
		return
	}
	for _, new := range news {
		go new.Pin()
	}

	// Unpin
	olds, err := getOldBuildsFromDevice(bh)
	if err != nil {
		return
	}
	for _, old := range olds {
		go old.Unpin()
	}
}

// Pin a build to the local IPFS cluster.
func (bh BuildHash) Pin() {
	const speed = 10 * 1024 * 1024
	const seconds = 60
	size := bh.Build.Size

	min := strconv.Itoa(1)
	max := strconv.Itoa(size/(speed*seconds) + 1)

	out, err := exec.Command("ipfs-cluster-ctl", "pin", "add", bh.IPFS, "--name", bh.Build.Filename, "--replication-min", min, "--replication-max", max).Output()
	if err != nil {
		fmt.Println("Failed to pin a build.")
		fmt.Println(aurora.Bold("Command :"), "ipfs-cluster-ctl", "pin", "add", aurora.Cyan(bh.IPFS), "--name", aurora.Green(bh.Build.Filename), "--replication-min", min, "--replication-max", max)

		// Log the error from the command
		ee, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println(string(ee.Stderr))
		}
	}

	fmt.Println(string(out))
}

// Unpin a build from the local IPFS cluster.
func (bh BuildHash) Unpin() {
	out, err := exec.Command("ipfs-cluster-ctl", "pin", "rm", bh.IPFS).Output()
	if err != nil {
		fmt.Println("Failed to unpin a build from the cluster.")
		fmt.Println(aurora.Bold("Command :"), "ipfs-cluster-ctl", "pin", "rm", aurora.Cyan(bh.IPFS))

		// Log the error from the command
		ee, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println(string(ee.Stderr))
		}
	}
	fmt.Println(string(out))
}
