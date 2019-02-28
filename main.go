package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"sync"
	"time"

	"github.com/logrusorgru/aurora"
)

func main() {

	// License
	fmt.Println("")
	fmt.Println("LOSGoI : LineageOS Goes to IPFS.")
	fmt.Println("Copyright © 2019 Nato Boram")
	fmt.Println("This program is free software : you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY ; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/.")
	fmt.Println("Contact : https://gitlab.com/NatoBoram/LOSGoI")
	fmt.Println("")

	// Check for IPFS
	ipfsCmdPath, err := exec.LookPath(ipfsCmdName)
	if err != nil {
		fmt.Println("IPFS is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("IPFS :", ipfsCmdPath)

	// Download device list
	devices, err := getDevices()
	if err != nil {
		fmt.Println("Couldn't get a list of devices.")
		fmt.Println(err.Error())
		return
	}

	// Count devices and builds
	fmt.Println("Received", aurora.Bold(len(devices)), "devices.")
	buildcount := countBuilds(devices)
	fmt.Println("Received", aurora.Bold(buildcount), "builds.")

	// Set an amount of workers
	// concurrency := runtime.NumCPU()
	concurrency := 1

	// Create channels
	jobs := make(chan Build, concurrency)
	results := make(chan HashedBuild)
	mtx := sync.Mutex{}

	// Create workers
	for w := 1; w <= concurrency; w++ {
		go urlstore(w, jobs, results)
	}

	mtx.Lock()
	go func() {

		// Receive results from the workers
		for i := 0; i < buildcount; i++ {
			hashedBuild := <-results
			fmt.Println(i+1, "/", buildcount, "|", hashedBuild.Hash, "|", hashedBuild.Build.Filename)
		}

		time.Sleep(time.Second)
		mtx.Unlock()
	}()

	// Create queue
	for device, builds := range devices {

		// Add a device to the queue
		fmt.Println("Adding device", aurora.Bold(device), "to the queue.")
		for _, build := range builds {

			// Add a build to the queue
			fmt.Println("Adding build", aurora.Green(build.Filename), "to the queue.")
			jobs <- build
		}
	}
	close(jobs)
	mtx.Lock() //jobs are done and results recieved

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

	return
}

func countBuilds(devices Devices) (count int) {
	for _, builds := range devices {
		for range builds {
			count++
		}
	}
	return
}

func urlstore(id int, builds <-chan Build, hash chan<- HashedBuild) {

	for build := range builds {

		// Log
		fmt.Println("Worker", aurora.Bold(id), "is processing the build", aurora.Green(build.Filename).String()+".")

		// URL
		filepath := mirrorbits + build.Filepath

		// Command
		out, err := exec.Command(ipfsCmdName, "urlstore", "add", "-t", filepath).Output()
		if err != nil {
			fmt.Println("Worker", aurora.Bold(id), "failed to execute the command.")
			fmt.Println("Command :", ipfsCmdName, "urlstore", "add", "-t", filepath)
			fmt.Println(err.Error())
			return
		}

		// out := "Test"

		// Log
		fmt.Println("Worker", aurora.Bold(id), "finished processing the build", aurora.Green(build.Filename).String()+".")

		// Return
		hash <- HashedBuild{
			Worker: id,
			Build:  &build,
			Hash:   string(out),
		}
	}
}
