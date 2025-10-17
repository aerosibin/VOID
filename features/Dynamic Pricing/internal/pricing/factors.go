package pricing

import "math/rand"

type Factors struct {
	TrafficDensity float64
	DistanceKm     float64
	DemandSurge    float64
}

func LoadFactors() *Factors {
	return &Factors{
		TrafficDensity: rand.Float64(),
		DistanceKm:     1 + rand.Float64()*10,
		DemandSurge:    0.5 + rand.Float64()*2,
	}
}
