// Package daysteps parses and summarises daily walking activity.
package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps stores a walk and the embedded user profile.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse reads steps and duration, committing both fields only on success.
func (ds *DaySteps) Parse(datastring string) error {
	fields := strings.Split(datastring, ",")
	if len(fields) != 2 {
		return fmt.Errorf("ожидаются два поля: шаги,продолжительность")
	}
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	duration, err := time.ParseDuration(fields[1])
	if err != nil {
		return fmt.Errorf("продолжительность прогулки: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	ds.Steps, ds.Duration = steps, duration
	return nil
}

// ActionInfo returns the walking distance and energy or an empty string on error.
func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, spentenergy.Distance(ds.Steps, ds.Height), calories), nil
}
