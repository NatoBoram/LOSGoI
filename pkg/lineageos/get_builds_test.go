package lineageos_test

import (
	"net/http"
	"testing"

	"gitlab.com/NatoBoram/LOSGoI/internal/self"
	"gitlab.com/NatoBoram/LOSGoI/pkg/lineageos"
)

func TestGet(t *testing.T) {
	service := lineageos.New(self.Agent(), &http.Client{})

	_, err := service.GetBuilds()
	if err != nil {
		t.Error("Couldn't get builds.", err)
	}
}
