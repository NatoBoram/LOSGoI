package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/logrusorgru/aurora"
)

func main() {

	// License
	fmt.Println("")
	fmt.Println(aurora.Bold("LOSGoI :"), "LineageOS Goes to IPFS.")
	fmt.Println("Copyright © 2019 Nato Boram")
	fmt.Println("This program is free software : you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY ; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/.")
	fmt.Println(aurora.Bold("Contact :"), aurora.Blue("https://gitlab.com/NatoBoram/LOSGoI"))
	fmt.Println("")

	// IPFS
	err := initIPFS()
	if err != nil {
		return
	}

	// IPFS Cluster
	err = initIPFSCluster()
	if err != nil {
		return
	}

	// Database
	err = initDatabase()
	if err != nil {
		return
	}
	defer db.Close()

	for {
		work()
		unpin()
		pin()
		time.Sleep(time.Hour)
	}
}

func initIPFS() (err error) {

	// Check for IPFS
	ipfsCmdPath, err := exec.LookPath("ipfs")
	if err != nil {
		fmt.Println("IPFS is not installed.")
		fmt.Println(err.Error())
		return
	}

	// Log
	fmt.Println(aurora.Bold("IPFS :"), aurora.Blue(ipfsCmdPath))

	// Use experimental features
	exec.Command("ipfs", "config", "--json", "Experimental.FilestoreEnabled", "true").Run()
	exec.Command("ipfs", "config", "--json", "Experimental.Libp2pStreamMounting", "true").Run()
	exec.Command("ipfs", "config", "--json", "Experimental.P2pHttpProxy", "true").Run()
	exec.Command("ipfs", "config", "--json", "Experimental.QUIC", "true").Run()
	exec.Command("ipfs", "config", "--json", "Experimental.ShardingEnabled", "true").Run()
	exec.Command("ipfs", "config", "--json", "Experimental.UrlstoreEnabled", "true").Run()

	return
}

func initIPFSCluster() (err error) {

	// Check for IPFS
	ipfsCmdPath, err := exec.LookPath("ipfs-cluster-ctl")
	if err != nil {
		fmt.Println("IPFS Cluster is not installed.")
		fmt.Println(err.Error())
		return
	}

	// Log
	fmt.Println(aurora.Bold("IPFS-Cluster :"), aurora.Blue(ipfsCmdPath))

	return
}

func initDatabase() (err error) {

	// Read the database config
	var database Database
	err = readDatabase(&database)
	if err != nil {
		fmt.Println("Could not load the database configuration.")
		fmt.Println(err.Error())
		writeTemplateDatabase()
		return
	}

	// Check for empty JSON
	if database.IsEmpty() {
		err = errors.New("Configuration is missing inside " + databasePath)
		fmt.Println(err.Error())
		return
	}

	// Connect to MariaDB
	db, err = sql.Open("mysql", database.ToConnectionString())
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

func work() {

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

	bhs, err := getLatestBuilds()
	if err != nil {
		return
	}

	for _, bh := range bhs {
		bh.Pin()
	}
}

func unpin() {

	bhs, err := getOldBuilds()
	if err != nil {
		return
	}

	for _, bh := range bhs {
		bh.Unpin()
	}
}
