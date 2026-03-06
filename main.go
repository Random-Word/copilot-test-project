package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Random-Word/copilot-test-project/conversion"
)

func main() {
	units := flag.String("units", "celsius", "temperature unit: celsius (c), fahrenheit (f), or kelvin (k)")
	flag.Parse()

	unit, err := conversion.ParseUnit(*units)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Demo: convert a sample temperature (20°C) to the requested unit
	sample := conversion.Convert(20, conversion.Celsius, unit)
	fmt.Printf("Current temperature: %s\n", conversion.FormatTemp(sample, unit))
}
