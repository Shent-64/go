package mascot_test

import (
	"testing"

	"example.com/go-learn-1/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gepher" {
		t.Fatal("Wrong mascot :(")
	}
}
