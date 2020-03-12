package math

import (
	"github.com/GDG-Cloud-Innopolis/Go-begginners/l3/test"
)

// Calc is a function that return an operation function
func Calc(operation test.CalcOperation) func(a, b float64) float64 {
	// Place your solution here
	switch operation {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	default:
		return func(a, b float64) float64 { return 0 }
	}
}


