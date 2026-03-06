package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const openWeatherMapBaseURL = "https://api.openweathermap.org/data/2.5/weather"

// WeatherCondition represents a weather condition from the API response.
type WeatherCondition struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

// MainWeather holds the primary weather measurements.
type MainWeather struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Humidity  int     `json:"humidity"`
}

// WeatherResponse represents the top-level API response.
type WeatherResponse struct {
	Name    string             `json:"name"`
	Weather []WeatherCondition `json:"weather"`
	Main    MainWeather        `json:"main"`
}

// APIError represents an error response from the API.
type APIError struct {
	Cod     interface{} `json:"cod"`
	Message string      `json:"message"`
}

func fetchWeather(city, apiKey string) (*WeatherResponse, error) {
	return fetchWeatherFromURL(openWeatherMapBaseURL, city, apiKey)
}

func fetchWeatherFromURL(baseURL, city, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", baseURL, city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		if json.Unmarshal(body, &apiErr) == nil && apiErr.Message != "" {
			return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, apiErr.Message)
		}
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	return &weather, nil
}

func displayWeather(w *WeatherResponse) {
	fmt.Printf("Weather for %s:\n", w.Name)
	fmt.Printf("  Temperature: %.1f°C (feels like %.1f°C)\n", w.Main.Temp, w.Main.FeelsLike)
	fmt.Printf("  Humidity:    %d%%\n", w.Main.Humidity)

	if len(w.Weather) > 0 {
		conditions := make([]string, len(w.Weather))
		for i, c := range w.Weather {
			conditions[i] = c.Description
		}
		fmt.Printf("  Conditions:  %s\n", strings.Join(conditions, ", "))
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: weather <city>")
	}

	city := strings.Join(os.Args[1:], " ")

	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPENWEATHERMAP_API_KEY environment variable is required")
	}

	weather, err := fetchWeather(city, apiKey)
	if err != nil {
		return err
	}

	displayWeather(weather)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
