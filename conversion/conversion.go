// Package conversion provides temperature unit conversion between Celsius, Fahrenheit, and Kelvin.
package conversion

import (
	"fmt"
	"strings"
)

// Unit represents a temperature unit.
type Unit string

const (
	Celsius    Unit = "celsius"
	Fahrenheit Unit = "fahrenheit"
	Kelvin     Unit = "kelvin"
)

// ParseUnit parses a string into a Unit, accepting common abbreviations.
func ParseUnit(s string) (Unit, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "c", "celsius":
		return Celsius, nil
	case "f", "fahrenheit":
		return Fahrenheit, nil
	case "k", "kelvin":
		return Kelvin, nil
	default:
		return "", fmt.Errorf("unknown unit %q: expected celsius (c), fahrenheit (f), or kelvin (k)", s)
	}
}

// Convert transforms a temperature value from one unit to another.
func Convert(value float64, from, to Unit) float64 {
	if from == to {
		return value
	}
	// Normalize to Celsius first
	c := toCelsius(value, from)
	return fromCelsius(c, to)
}

func toCelsius(value float64, from Unit) float64 {
	switch from {
	case Fahrenheit:
		return (value - 32) * 5 / 9
	case Kelvin:
		return value - 273.15
	default:
		return value
	}
}

func fromCelsius(c float64, to Unit) float64 {
	switch to {
	case Fahrenheit:
		return c*9/5 + 32
	case Kelvin:
		return c + 273.15
	default:
		return c
	}
}

// FormatTemp formats a temperature value with its unit symbol.
func FormatTemp(value float64, unit Unit) string {
	switch unit {
	case Fahrenheit:
		return fmt.Sprintf("%.1f°F", value)
	case Kelvin:
		return fmt.Sprintf("%.1f K", value)
	default:
		return fmt.Sprintf("%.1f°C", value)
	}
}
