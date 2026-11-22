package daysteps

import (
	"testing"
	"time"
)

func TestParsePreservesDayStepsOnError(t *testing.T) {
	original := DaySteps{Steps: 1000, Duration: time.Hour}
	for _, input := range []string{"broken", "1500,invalid", "0,1h", "1500,0m"} {
		walk := original
		if walk.Parse(input) == nil {
			t.Fatalf("expected error for %q", input)
		}
		if walk != original {
			t.Fatalf("partial state change for %q: %+v", input, walk)
		}
	}
}
