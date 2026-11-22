// Package personaldata contains the user's profile.
package personaldata

import "fmt"

// Personal stores the user's name, weight in kilograms and height in metres.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print writes the profile to standard output with two decimal places.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
