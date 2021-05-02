package space

type Planet string

func Age(age float64, planet Planet) float64 {
	earthSeconds := 31557600.0
	switch planet {
	case "Mercury":
		return age / (0.2408467 * earthSeconds)
	case "Venus":
		return age / (0.61519726 * earthSeconds)
	case "Earth":
		return age / 31557600
	case "Mars":
		return age / (1.8808158 * earthSeconds)
	case "Jupiter":
		return age / (11.862615 * earthSeconds)
	case "Saturn":
		return age / (29.447498 * earthSeconds)
	case "Uranus":
		return age / (84.016846 * earthSeconds)
	case "Neptune":
		return age / (164.79132 * earthSeconds)
	default:
		return 1
	}
}
