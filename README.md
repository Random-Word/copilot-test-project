# Weather CLI

A simple command-line weather tool written in Go.

## Features
- [x] Colored terminal output (blue ≤10°C, yellow ≤25°C, red >25°C)
- [ ] Fetch current weather by city name
- [ ] 5-day forecast display
- [ ] Temperature unit conversion (C/F/K)
- [ ] Cache recent lookups to reduce API calls

## Usage
```
weather london
weather --no-color tokyo
weather --forecast 5 tokyo
weather --units fahrenheit paris
```

### Flags
| Flag | Description |
|------|-------------|
| `--no-color` | Disable colored output |
