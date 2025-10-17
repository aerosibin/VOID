package pricing

import "VOID/config"

// final = base + distance*distance_weight + traffic*traffic_weight + demand*demand_weight
func CalculatePrice(base, distanceKm, trafficScore, demandIndex float64, pcfg config.PricingCfg) float64 {
	return base + distanceKm*pcfg.DistanceWeight + trafficScore*pcfg.TrafficWeight + demandIndex*pcfg.DemandWeight
}
