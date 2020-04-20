package losgoi

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
	PermPrivateDirectory os.FileMode = 0700
	PermPrivateFile      os.FileMode = 0600
)

// Special Hashes
const (
	hashEmptyFolder = "QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn"
)

// Calculation
const (
	speed   = 10 * 1024 * 1024
	seconds = 60
)

// Paralellization
const (
	coHashPin   = 4
	coHashUnpin = 4
)
