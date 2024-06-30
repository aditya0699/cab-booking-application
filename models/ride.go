package models

import "time"

type Ride struct {
	ID          string
	UserID      string
	DriverID    string
	CabId       string
	Pickup      *Location
	Drop        *Location
	CabType     string
	Status      string
	StartTime   time.Time
	EndTime     time.Time
	TotalAmount float64
}

type RideRequest struct {
	SrcLat  float64
	SrcLng  float64
	DestLat float64
	DestLng float64
	CabType string
}
