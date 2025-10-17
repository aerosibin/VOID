package pricing

type PricingRequest struct {
	OrderID      uint
	DistanceKm   float64
	TrafficScore float64
	DemandIndex  float64
}
