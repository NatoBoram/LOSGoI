package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {

	// Prevent overshadowing
	var err error

	// Check for IPFS
	ipfscmdpath, err = exec.LookPath("ipfs")
	if err != nil {
		fmt.Println("IPFS is not installed.")
		return
	}
	fmt.Println("IPFS :", ipfscmdpath)

	// Devices
	devices, err := getDevices()
	if err != nil {
		fmt.Println("Couldn't get a list of devices.")
		fmt.Println(err.Error())
		return
	}

	// Device
	for device, builds := range *devices {

		// go processDevice(&device, &builds)
		processDevice(&device, &builds)

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

func processDevice(device *string, builds *[]Build) {

	// Log
	fmt.Println("Processing device", *device+".")

	// Build
	for _, build := range *builds {

		// go processBuild(build)
		processBuild(build)

	}
}

func processBuild(build Build) {

	// Log
	fmt.Println("Processing build", build.FileName)

	filepath := "https://mirrorbits.lineageos.org" + build.FilePath

	// Prepare command
	out, err := exec.Command(ipfscmdpath, "urlstore add "+filepath, "-t").Output()
	if err != nil {
		fmt.Println("Couldn't execute the command.")
		fmt.Println("Command :", ipfscmdpath, "urlstore add "+filepath, "-t")
		fmt.Println(err.Error())
		return
	}

	// Log
	fmt.Println("Output : ", string(out))
}
