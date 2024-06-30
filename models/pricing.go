package models

type PricingTier struct {
	Tier     string
	BaseFare float64
	PerKM    float64
}
