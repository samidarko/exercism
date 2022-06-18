package meteorology

import (
	"fmt"
)

// TemperatureUnit type
type TemperatureUnit int

// String stringify TemperatureUnit
func (u TemperatureUnit) String() string {
	switch u {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	default:
		return ""
	}
}

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Temperature type
type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// String stringify Temperature
func (t *Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit.String())
}

// SpeedUnit type
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// String stringify SpeedUnit
func (u SpeedUnit) String() string {
	switch u {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	default:
		return ""
	}
}

// Speed type
type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// String stringify Speed
func (s *Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// String stringify MeteorologyData
func (d *MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", d.location, d.temperature.String(), d.windDirection, d.windSpeed.String(), d.humidity)
}
