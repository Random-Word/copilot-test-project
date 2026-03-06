package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Temperature thresholds in Celsius for color bands.
const (
	coldMax = 10.0
	mildMax = 25.0
)

// ColorPrinter formats weather output with temperature-based colors.
type ColorPrinter struct {
	cold *color.Color
	mild *color.Color
	hot  *color.Color
}

// NewColorPrinter returns a printer with blue/yellow/red temperature bands.
func NewColorPrinter() *ColorPrinter {
	return &ColorPrinter{
		cold: color.New(color.FgBlue, color.Bold),
		mild: color.New(color.FgYellow, color.Bold),
		hot:  color.New(color.FgRed, color.Bold),
	}
}

// FormatTemp returns a color-formatted temperature string.
func (cp *ColorPrinter) FormatTemp(celsius float64) string {
	label := fmt.Sprintf("%.1f°C", celsius)
	switch {
	case celsius <= coldMax:
		return cp.cold.Sprint(label)
	case celsius <= mildMax:
		return cp.mild.Sprint(label)
	default:
		return cp.hot.Sprint(label)
	}
}

// PrintWeather outputs a sample weather line with colored temperature.
func (cp *ColorPrinter) PrintWeather(city string, celsius float64, condition string) {
	fmt.Printf("%s: %s — %s\n", city, cp.FormatTemp(celsius), condition)
}

func main() {
	noColor := flag.Bool("no-color", false, "disable colored output")
	flag.Parse()

	if *noColor {
		color.NoColor = true
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: weather [--no-color] <city>")
		os.Exit(1)
	}

	city := args[0]
	printer := NewColorPrinter()

	// Placeholder data until real API integration is added.
	samples := map[string]struct {
		temp      float64
		condition string
	}{
		"london": {8.0, "Overcast"},
		"tokyo":  {22.0, "Partly cloudy"},
		"dubai":  {38.0, "Sunny"},
	}

	if data, ok := samples[city]; ok {
		printer.PrintWeather(city, data.temp, data.condition)
	} else {
		printer.PrintWeather(city, 15.0, "No data (placeholder)")
	}
}
