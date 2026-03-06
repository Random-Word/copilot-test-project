package conversion

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.01
}

func TestConvertIdentity(t *testing.T) {
	for _, u := range []Unit{Celsius, Fahrenheit, Kelvin} {
		if got := Convert(100, u, u); got != 100 {
			t.Errorf("Convert(100, %s, %s) = %f, want 100", u, u, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		c, f float64
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
		{37, 98.6},
	}
	for _, tt := range tests {
		got := Convert(tt.c, Celsius, Fahrenheit)
		if !almostEqual(got, tt.f) {
			t.Errorf("Convert(%f, C, F) = %f, want %f", tt.c, got, tt.f)
		}
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		f, c float64
	}{
		{32, 0},
		{212, 100},
		{-40, -40},
		{98.6, 37},
	}
	for _, tt := range tests {
		got := Convert(tt.f, Fahrenheit, Celsius)
		if !almostEqual(got, tt.c) {
			t.Errorf("Convert(%f, F, C) = %f, want %f", tt.f, got, tt.c)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		c, k float64
	}{
		{0, 273.15},
		{100, 373.15},
		{-273.15, 0},
	}
	for _, tt := range tests {
		got := Convert(tt.c, Celsius, Kelvin)
		if !almostEqual(got, tt.k) {
			t.Errorf("Convert(%f, C, K) = %f, want %f", tt.c, got, tt.k)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	got := Convert(0, Kelvin, Fahrenheit)
	want := -459.67
	if !almostEqual(got, want) {
		t.Errorf("Convert(0, K, F) = %f, want %f", got, want)
	}
}

func TestParseUnit(t *testing.T) {
	tests := []struct {
		input string
		want  Unit
		err   bool
	}{
		{"c", Celsius, false},
		{"celsius", Celsius, false},
		{"Celsius", Celsius, false},
		{"f", Fahrenheit, false},
		{"fahrenheit", Fahrenheit, false},
		{"k", Kelvin, false},
		{"kelvin", Kelvin, false},
		{"KELVIN", Kelvin, false},
		{"invalid", "", true},
	}
	for _, tt := range tests {
		got, err := ParseUnit(tt.input)
		if (err != nil) != tt.err {
			t.Errorf("ParseUnit(%q) error = %v, wantErr %v", tt.input, err, tt.err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseUnit(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestFormatTemp(t *testing.T) {
	tests := []struct {
		value float64
		unit  Unit
		want  string
	}{
		{20.0, Celsius, "20.0°C"},
		{68.0, Fahrenheit, "68.0°F"},
		{293.1, Kelvin, "293.1 K"},
	}
	for _, tt := range tests {
		got := FormatTemp(tt.value, tt.unit)
		if got != tt.want {
			t.Errorf("FormatTemp(%f, %s) = %q, want %q", tt.value, tt.unit, got, tt.want)
		}
	}
}
