package self_test

import (
	"testing"

	"gitlab.com/NatoBoram/LOSGoI/internal/self"
)

func TestAgent(t *testing.T) {
	agent := self.Agent()
	t.Logf("User-Agent: %s", agent)
}
