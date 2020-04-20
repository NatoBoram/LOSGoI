package losgoi

import (
	"time"

	"gitlab.com/NatoBoram/LOSGoI/pkg/lineageos"
)

// Logger defines what this program needs to be logged.
type Logger interface {
	Fetching() (now time.Time)
	Fetched(start time.Time, builds lineageos.Builds)

	Skipping(build Build)
}
