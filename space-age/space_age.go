package space

import "math"

type Planet string

var orbitalPeriods = map[string]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, p Planet) float64 {
	yearsInEarth := seconds / 60 / 60 / 24 / 365.25
	yearsInPlanet := yearsInEarth / orbitalPeriods[string(p)]
	return math.Round(yearsInPlanet*100) / 100
}
