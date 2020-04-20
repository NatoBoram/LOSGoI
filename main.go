package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/logrusorgru/aurora"
)

func main() {

	// License
	fmt.Println("")
	fmt.Println(aurora.Bold("LOSGoI :"), "LineageOS Goes to IPFS.")
	fmt.Println("Copyright © 2018-2020 Nato Boram")
	fmt.Println("This program is free software : you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY ; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/.")
	fmt.Println(aurora.Bold("Contact :"), aurora.Blue("https://gitlab.com/NatoBoram/LOSGoI"))
	fmt.Println("")

	// IPFS
	err := initIPFS()
	if err != nil {
		return
	}

	// Database
	err = initDatabase()
	if err != nil {
		return
	}
	defer db.Close()

	// Put LineageOS on IPFS
	for {

		// Health
		recover()
		pin()

		// Add new builds
		unpin()
		add()

		// Add lost builds
		unpin()
		readd()

		// Give some rest to LineageOS
		time.Sleep(time.Hour)
	}
}

func initIPFS() (err error) {

	// Check for IPFS
	path, err := exec.LookPath("ipfs")
	if err != nil {
		fmt.Println("IPFS is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("IPFS :"), aurora.Blue(path))

	// Enable sharding
	exec.Command("ipfs", "config", "--json", "Experimental.ShardingEnabled", "true").Run()

	// Check for IPFS Cluster Service
	path, err = exec.LookPath("ipfs-cluster-service")
	if err != nil {
		fmt.Println("IPFS Cluster Service is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("IPFS Cluster Service :"), aurora.Blue(path))

	// Check for IPFS Cluster Control
	path, err = exec.LookPath("ipfs-cluster-ctl")
	if err != nil {
		fmt.Println("IPFS Cluster Control is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("IPFS Cluster Control :"), aurora.Blue(path))

	fmt.Println("")
	return
}

func initDatabase() (err error) {

	// Read the database config
	var database Database
	err = database.read()
	if err != nil {
		fmt.Println("Could not load the database configuration.")
		fmt.Println(err.Error())
		terr := database.template()
		if terr != nil {
			fmt.Println("Could not write a database template.")
			fmt.Println(err.Error())
			return terr
		}
		return
	}

	// Check for empty JSON
	if database.IsEmpty() {
		err = errors.New("Configuration is missing inside " + databasePath)
		fmt.Println(err.Error())
		return
	}

	// Connect to MariaDB
	db, err = sql.Open("mysql", database.string())
	if err != nil {
		fmt.Println("Could not connect to the database.")
		fmt.Println(err.Error())
		return
	}

	// Test the connection to MariaDB
	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping the database.")
		fmt.Println(err.Error())
		return
	}

	return
}

func add() {
	fmt.Println("Adding new builds to the cluster...")

	// Download device list
	devices, err := getDevices()
	if err != nil {
		fmt.Println("Couldn't get a list of devices.")
		fmt.Println(err.Error())
		return
	}

	// Count devices and builds
	fmt.Println("Received", aurora.Bold(len(devices)), "devices.")
	buildcount := devices.Count()
	fmt.Println("Received", aurora.Bold(buildcount), "builds.")
	fmt.Println("")

	// Hash builds
	devices.Hash()
}

func getDevices() (devices Devices, err error) {

	// GET
	resp, err := http.Get(api)
	if err != nil {
		fmt.Println("Couldn't get a response from LineageOS.")
		return
	}
	defer resp.Body.Close()

	// JSON
	err = json.NewDecoder(resp.Body).Decode(&devices)
	if err != nil {
		fmt.Println("Couldn't decode the received JSON.")
		return
	}

	// Builds are received in a map.
	devices.Name()
	return
}

func pin() {
	fmt.Println("Pinning latest builds...")
	start := time.Now()

	bhs, err := getLatestBuilds()
	if err != nil {
		return
	}

	// Create queue
	bhc := make(chan *BuildHash, coHashPin)
	go func() {
		for _, bh := range bhs {
			bhc <- bh
		}
		close(bhc)
	}()

	var wg sync.WaitGroup
	wg.Add(len(bhs))

	// Pin asynchronously
	for c := 1; c <= coHashPin; c++ {
		go func() {
			for bh := range bhc {
				bh.Pin()
				wg.Done()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Pinned latest builds in", aurora.Bold(time.Since(start).String()).String()+".")
}

func unpin() {
	fmt.Println("Unpinning old builds...")
	start := time.Now()

	bhs, err := getOldBuilds()
	if err != nil {
		return
	}

	// Create queue
	bhc := make(chan *BuildHash, coHashUnpin)
	go func() {
		for _, bh := range bhs {
			bhc <- bh
		}
		close(bhc)
	}()

	var wg sync.WaitGroup
	wg.Add(len(bhs))

	// Unpin asynchronously
	for c := 1; c <= coHashUnpin; c++ {
		go func() {
			for bh := range bhc {
				bh.Unpin()
				wg.Done()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Unpinned old builds in", aurora.Bold(time.Since(start).String()).String()+".")

	gc()
}

func recover() {
	ctlsync()

	_, err := exec.Command("ipfs-cluster-ctl", "recover", "--local").Output()
	if err != nil {
		fmt.Println("Failed to recover.")
		fmt.Println(aurora.Bold("Command :"), "ipfs-cluster-ctl", "recover", "--local")

		// Log the error from the command
		ee, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println(string(ee.Stderr))
		}
	}
}

func gc() {
	fmt.Println("Starting garbage collection...")
	start := time.Now()

	out, err := exec.Command("ipfs", "repo", "gc").Output()
	if err != nil {
		fmt.Println("Failed to recover.")
		fmt.Println(aurora.Bold("Command :"), "ipfs", "repo", "gc")

		// Log the error from the command
		ee, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println(string(ee.Stderr))
		}
	}

	fmt.Println(string(out))
	fmt.Println("Garbage collected in", aurora.Bold(time.Since(start).String()).String()+".")
}

func eta(start time.Time, index int, total int) (left time.Duration) {
	return time.Since(start) / (time.Duration(index) * time.Millisecond) * (time.Duration(total-index) * time.Millisecond)
}

func readd() {
	fmt.Println("Recovering lost builds...")
	start := time.Now()

	bhs, err := getLatestBuilds()
	if err != nil {
		return
	}

	var builds []*Build
	for _, bh := range bhs {

		// Get this build's status
		out, err := exec.Command("ipfs-cluster-ctl", "status", bh.IPFS).Output()
		if err != nil {
			return
		}

		// Check for not "PINNED"
		if !strings.Contains(string(out), "PINNED") {

			// Cluster lost this build
			builds = append(builds, bh.Build)
		}
	}

	total := len(builds)
	for index, build := range builds {

		build.Hash(index, total)

		// Estimated Time Left
		fmt.Println("Estimated Time Left :", aurora.Bold(eta(start, index+1, total)).String()+".")
	}

	// Duration
	fmt.Println("Hashed in", aurora.Bold(time.Since(start).String()).String()+".")
}

func ctlsync() {
	_, err := exec.Command("ipfs-cluster-ctl", "sync").Output()
	if err != nil {
		fmt.Println("Failed to sync.")
		fmt.Println(aurora.Bold("Command :"), "ipfs-cluster-ctl", "sync")

		// Log the error from the command
		ee, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println(string(ee.Stderr))
		}
	}
}

func percent(index int, total int) string {
	return fmt.Sprintf("%3.2f%%", float64(index)/float64(total)*100)
}
