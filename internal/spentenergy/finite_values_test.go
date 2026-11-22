package spentenergy

import (
	"math"
	"testing"
	"time"
)

func TestInvalidAndNonFiniteValues(t *testing.T) {
	for _, value := range []float64{0, -1, math.NaN(), math.Inf(1), math.Inf(-1)} {
		if Distance(100, value) != 0 {
			t.Errorf("Distance accepted %v", value)
		}
		if calories, err := RunningSpentCalories(100, value, 1.8, time.Hour); err == nil || calories != 0 {
			t.Errorf("running accepted weight %v", value)
		}
		if calories, err := WalkingSpentCalories(100, 80, value, time.Hour); err == nil || calories != 0 {
			t.Errorf("walking accepted height %v", value)
		}
	}
}
