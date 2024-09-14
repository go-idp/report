package report

import (
	"testing"
)

func TestReport(t *testing.T) {
	if err := Report("self", "test report in ci", nil); err != nil {
		t.Fatal(err)
	}
}
