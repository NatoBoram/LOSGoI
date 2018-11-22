package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/logrusorgru/aurora"
)

func main() {
	// workertest()

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
	concurrency := runtime.NumCPU()
	// concurrency := 1

	// Create channels
	jobs := make(chan Build, concurrency)
	results := make(chan HashedBuild)

	// Create workers
	for w := 1; w <= concurrency; w++ {
		go urlstore(w, jobs, results)
	}

	// Create queue
	for device, builds := range devices {

		// Add a device to the queue
		fmt.Println("Adding device", aurora.Bold(device), "to the queue.")
		for _, build := range builds {

			// Add a build to the queue
			fmt.Println("Adding build", aurora.Bold(build.FileName), "to the queue.")
			jobs <- build
		}
	}
	close(jobs)

	// Receive results from the workers
	for i := 0; i < buildcount; i++ {
		hashedBuild := <-results
		fmt.Println(i, "/", buildcount, "|", hashedBuild.Hash, "|", hashedBuild.Build.FileName)
	}
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
		fmt.Println("Worker", aurora.Bold(id), "is processing the build", aurora.Bold(build.FileName).String()+".")
		/*
			// URL
			filepath := mirrorbits + build.FilePath

			// Command
			out, err := exec.Command(ipfsCmdName, "urlstore", "add", filepath, "-t").Output()
			if err != nil {
				fmt.Println("Couldn't execute the command.")
				fmt.Println("Command :", ipfsCmdName, "urlstore", "add", filepath, "-t")
				fmt.Println(err.Error())
				return
			}
		*/
		out := "Test"

		// Log
		fmt.Println("Worker", aurora.Bold(id), "finished processing the build", aurora.Bold(build.FileName).String()+".")

		// Return
		hash <- HashedBuild{
			Worker: id,
			Build:  &build,
			Hash:   string(out),
		}
	}
}
