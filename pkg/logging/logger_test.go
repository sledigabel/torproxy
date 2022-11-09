package logging

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	if _, err := GetLogger(); err != nil {
		t.Fatalf("Logger couldn't be initialised: %s", err)
	}
}
