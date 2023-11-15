package api_test

import (
	"testing"

	"miqgo.com/go/cmasters/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
