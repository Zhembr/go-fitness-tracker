// Package actioninfo processes any activity implementing DataParser.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser parses an activity and produces its human-readable description.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info logs invalid records and continues processing the rest of the dataset.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("парсинг %q: %v", data, err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("расчёт активности %q: %v", data, err)
			continue
		}
		fmt.Println(info)
	}
}
