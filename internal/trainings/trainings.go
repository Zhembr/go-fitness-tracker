// Package trainings parses workouts and formats their calculated results.
package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	running = "Бег"
	walking = "Ходьба"
)

// Training is a workout with an embedded user profile.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse reads steps, activity type and duration without changing state on error.
// Unknown activity types are reported by ActionInfo, not by the parser.
func (t *Training) Parse(datastring string) error {
	fields := strings.Split(datastring, ",")
	if len(fields) != 3 {
		return fmt.Errorf("ожидаются три поля: шаги,тип,продолжительность")
	}
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	duration, err := time.ParseDuration(fields[2])
	if err != nil {
		return fmt.Errorf("продолжительность тренировки: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	t.Steps, t.TrainingType, t.Duration = steps, fields[1], duration
	return nil
}

// ActionInfo calculates a supported workout and returns its formatted summary.
func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error
	switch t.TrainingType {
	case running:
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case walking:
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Height),
		spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration), calories), nil
}
