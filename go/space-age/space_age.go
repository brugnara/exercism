package space

// Planet holds planet name
type Planet string

func (p Planet) earth(x float64) float64 {
	return 31557600 * x
}

func (p Planet) ToPeriod() float64 {
	switch p {
	case "Mercury":
		return p.earth(0.2408467)
	case "Venus":
		return p.earth(0.61519726)
	case "Earth":
		return p.earth(1)
	case "Mars":
		return p.earth(1.8808158)
	case "Jupiter":
		return p.earth(11.862615)
	case "Saturn":
		return p.earth(29.447498)
	case "Uranus":
		return p.earth(84.016846)
	case "Neptune":
		return p.earth(164.79132)
	}
	return 0
}

// Age computes age based on livin Planet
func Age(t float64, p Planet) float64 {
	return t / p.ToPeriod()
}
