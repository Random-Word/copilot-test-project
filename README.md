# Weather CLI

A simple command-line weather tool written in Go.

## Features
- [x] Fetch current weather by city name (temperature, humidity, conditions)
- [ ] 5-day forecast display
- [ ] Temperature unit conversion (C/F/K)
- [ ] Cache recent lookups to reduce API calls
- [ ] Colored terminal output

## Setup

Set your OpenWeatherMap API key:
```
export OPENWEATHERMAP_API_KEY="your-api-key"
```

## Usage
```
weather london
weather new york
weather --forecast 5 tokyo
weather --units fahrenheit paris
```
