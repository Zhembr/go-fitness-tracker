package trainings

import (
	"testing"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

func TestParsePreservesTrainingOnError(t *testing.T) {
	original := Training{Steps: 1000, TrainingType: "Бег", Duration: time.Hour, Personal: personaldata.Personal{Name: "Тест", Weight: 75, Height: 1.75}}
	for _, input := range []string{"broken", "1500,Ходьба,invalid", "0,Бег,1h", "1500,Бег,0m"} {
		training := original
		if training.Parse(input) == nil {
			t.Fatalf("expected error for %q", input)
		}
		if training != original {
			t.Fatalf("partial state change for %q: %+v", input, training)
		}
	}
}
