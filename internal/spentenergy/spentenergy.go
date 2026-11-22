// Package spentenergy calculates distance, speed and energy spent during activity.
package spentenergy

import (
	"fmt"
	"math"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// Distance returns the distance in kilometres for a height measured in metres.
// Invalid step counts and heights produce zero distance.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || !positiveFinite(height) {
		return 0
	}
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

// MeanSpeed returns kilometres per hour, or zero for invalid input.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// RunningSpentCalories returns the energy spent running, validating all inputs.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	}
	if !positiveFinite(weight) {
		return 0, fmt.Errorf("вес должен быть положительным конечным числом: %v", weight)
	}
	if !positiveFinite(height) {
		return 0, fmt.Errorf("рост должен быть положительным конечным числом: %v", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной: %v", duration)
	}
	calories := weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH
	if math.IsInf(calories, 0) || math.IsNaN(calories) {
		return 0, fmt.Errorf("результат расчёта калорий вне допустимого диапазона")
	}
	return calories, nil
}

// WalkingSpentCalories applies the walking correction to running energy.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}

func positiveFinite(value float64) bool {
	return value > 0 && !math.IsInf(value, 0) && !math.IsNaN(value)
}
