package lineageos_test

import (
	"net/http"
	"testing"

	"gitlab.com/NatoBoram/LOSGoI/internal/self"
	"gitlab.com/NatoBoram/LOSGoI/pkg/lineageos"
)

func TestNew(t *testing.T) {
	service := lineageos.New(self.Agent(), &http.Client{})
	if service == nil {
		t.Error("The service could not be created.")
	}
}
