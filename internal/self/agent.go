package self

import (
	"fmt"
)

// Agent generates a user agent.
func Agent() string {
	return fmt.Sprintf("%s/%s (%s; %s; +%s) %s", User, Version, OS, Arch, Contact, Go)
}
