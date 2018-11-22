package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {

	// Check for IPFS
	ipfsCmdPath, err := exec.LookPath(ipfsCmdName)
	if err != nil {
		fmt.Println("IPFS is not installed.")
		return
	}
	fmt.Println("IPFS :", ipfsCmdPath)

	// Devices
	devices, err := getDevices()
	if err != nil {
		fmt.Println("Couldn't get a list of devices.")
		fmt.Println(err.Error())
		return
	}

	// Channels
	jobs := make(chan Build)
	results := make(chan string)

	// Workers
	for w := 1; w <= runtime.NumCPU(); w++ {
		go urlstore(jobs, results)
	}

	// Device
	for device, builds := range *devices {

		// Log
		fmt.Println("Adding device", device, "to the queue.")

		// Build
		for _, build := range builds {

			// Log
			fmt.Println("Adding build", build.FileName, "to the queue.")

			jobs <- build
		}
	}
	close(jobs)

	// Hash
	for hash := range results {
		fmt.Println(hash)
	}
}

func getDevices() (devices *Devices, err error) {

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

	// Log
	fmt.Println("Received", len(*devices), "devices.")
	return
}

func urlstore(builds <-chan Build, hash chan<- string) {

	for build := range builds {

		// Log
		fmt.Println("Processing build", build.FileName)

		// URL
		filepath := "https://mirrorbits.lineageos.org" + build.FilePath

		// Command
		out, err := exec.Command(ipfsCmdName, "urlstore", "add", filepath, "-t").Output()
		if err != nil {
			fmt.Println("Couldn't execute the command.")
			fmt.Println("Command :", ipfsCmdName, "urlstore", "add", filepath, "-t")
			fmt.Println(err.Error())
			return
		}

		hash <- string(out)
	}
}
