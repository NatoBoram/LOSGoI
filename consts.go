package main

import "os"

// LineageOS
const mirrorbits = "https://mirrorbits.lineageos.org"
const api = mirrorbits + "/api/v1/builds/"

// Paths
const (
	rootFolder   = "./LOSGoI"
	databasePath = rootFolder + "/database.json"
)

// Permissions
const (
	permPrivateDirectory os.FileMode = 0700
	permPrivateFile      os.FileMode = 0600
)
