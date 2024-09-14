package report

import (
	"testing"
	"time"
)

func TestReport(t *testing.T) {
	if err := Report("self", "test report in ci", nil); err != nil {
		t.Fatal(err)
	}
	Report("self", "test report in ci 2", nil)
	time.Sleep(31 * time.Second)
	Report("self", "test report in ci 3", nil)
}
