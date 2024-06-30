package models

type Cab struct {
	ID           string
	Type         string
	PricingTier  string
	LicensePlate string
	Location     *Location
	Driver       *Driver
	Speed        float64
}
