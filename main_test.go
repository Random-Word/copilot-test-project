package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fatih/color"
)

func TestFetchWeather_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("q") != "London" {
			t.Errorf("expected city=London, got %s", q.Get("q"))
		}
		if q.Get("units") != "metric" {
			t.Errorf("expected units=metric, got %s", q.Get("units"))
		}

		resp := WeatherResponse{
			Name: "London",
			Weather: []WeatherCondition{
				{ID: 800, Main: "Clear", Description: "clear sky"},
			},
			Main: MainWeather{
				Temp:      15.2,
				FeelsLike: 14.0,
				TempMin:   13.0,
				TempMax:   17.0,
				Humidity:  72,
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Temporarily override the base URL by calling the server directly
	weather, err := fetchWeatherFromURL(server.URL, "London", "test-key")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if weather.Name != "London" {
		t.Errorf("expected name=London, got %s", weather.Name)
	}
	if weather.Main.Temp != 15.2 {
		t.Errorf("expected temp=15.2, got %f", weather.Main.Temp)
	}
	if weather.Main.Humidity != 72 {
		t.Errorf("expected humidity=72, got %d", weather.Main.Humidity)
	}
	if len(weather.Weather) != 1 || weather.Weather[0].Description != "clear sky" {
		t.Errorf("unexpected conditions: %v", weather.Weather)
	}
}

func TestFetchWeather_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(APIError{
			Cod:     "404",
			Message: "city not found",
		})
	}))
	defer server.Close()

	_, err := fetchWeatherFromURL(server.URL, "Nonexistent", "test-key")
	if err == nil {
		t.Fatal("expected error for non-existent city")
	}
	if got := err.Error(); got != "API error (404): city not found" {
		t.Errorf("unexpected error message: %s", got)
	}
}

func TestFetchWeather_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not json"))
	}))
	defer server.Close()

	_, err := fetchWeatherFromURL(server.URL, "London", "test-key")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestTempColor(t *testing.T) {
	tests := []struct {
		temp     float64
		wantAttr color.Attribute
	}{
		{-5, color.FgBlue},
		{0, color.FgBlue},
		{9.9, color.FgBlue},
		{10, color.FgYellow},
		{20, color.FgYellow},
		{25, color.FgYellow},
		{25.1, color.FgRed},
		{40, color.FgRed},
	}

	for _, tt := range tests {
		c := tempColor(tt.temp)
		want := color.New(tt.wantAttr)
		if c.Sprint("x") != want.Sprint("x") {
			t.Errorf("tempColor(%.1f): got different color than expected %v", tt.temp, tt.wantAttr)
		}
	}
}

func TestDisplayWeather_NoColor(t *testing.T) {
	color.NoColor = true
	defer func() { color.NoColor = false }()

	w := &WeatherResponse{
		Name: "TestCity",
		Weather: []WeatherCondition{
			{ID: 800, Main: "Clear", Description: "clear sky"},
		},
		Main: MainWeather{
			Temp:      5.0,
			FeelsLike: 3.0,
			Humidity:  50,
		},
	}
	// Should not panic when colors are disabled
	displayWeather(w)
}
