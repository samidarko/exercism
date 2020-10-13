package space

type Planet string

const earthOrbitalPeriod = 31557600 // in seconds

func Age(seconds float64, planet Planet) float64 {

	switch planet {
	case "Mercury":
		return seconds / (earthOrbitalPeriod * 0.2408467)
	case "Venus":
		return seconds / (earthOrbitalPeriod * 0.61519726)
	case "Earth":
		return seconds / earthOrbitalPeriod
	case "Mars":
		return seconds / (earthOrbitalPeriod * 1.8808158)
	case "Jupiter":
		return seconds / (earthOrbitalPeriod * 11.862615)
	case "Saturn":
		return seconds / (earthOrbitalPeriod * 29.447498)
	case "Uranus":
		return seconds / (earthOrbitalPeriod * 84.016846)
	case "Neptune":
		return seconds / (earthOrbitalPeriod * 164.79132)
	default:
		return 0
	}

}
