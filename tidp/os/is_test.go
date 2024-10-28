package os

import "testing"

func TestIsMacOS(t *testing.T) {
	if !IsMacOS() {
		t.Error("Expected IsMacOS() to return true")
	}
}
