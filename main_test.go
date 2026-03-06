package main

import (
	"testing"

	"github.com/fatih/color"
)

func init() {
	// Disable color in tests for deterministic output.
	color.NoColor = true
}

func TestFormatTemp_Cold(t *testing.T) {
	p := NewColorPrinter()
	got := p.FormatTemp(-5.0)
	want := "-5.0°C"
	if got != want {
		t.Errorf("FormatTemp(-5.0) = %q, want %q", got, want)
	}
}

func TestFormatTemp_ColdBoundary(t *testing.T) {
	p := NewColorPrinter()
	got := p.FormatTemp(10.0)
	want := "10.0°C"
	if got != want {
		t.Errorf("FormatTemp(10.0) = %q, want %q", got, want)
	}
}

func TestFormatTemp_Mild(t *testing.T) {
	p := NewColorPrinter()
	got := p.FormatTemp(20.0)
	want := "20.0°C"
	if got != want {
		t.Errorf("FormatTemp(20.0) = %q, want %q", got, want)
	}
}

func TestFormatTemp_MildBoundary(t *testing.T) {
	p := NewColorPrinter()
	got := p.FormatTemp(25.0)
	want := "25.0°C"
	if got != want {
		t.Errorf("FormatTemp(25.0) = %q, want %q", got, want)
	}
}

func TestFormatTemp_Hot(t *testing.T) {
	p := NewColorPrinter()
	got := p.FormatTemp(35.0)
	want := "35.0°C"
	if got != want {
		t.Errorf("FormatTemp(35.0) = %q, want %q", got, want)
	}
}
