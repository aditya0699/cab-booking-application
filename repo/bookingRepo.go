package repo

import (
	"time"

	"github.com/cab-booking-application/models"
)

type IBookingRepo interface {
	GetAllCabsByLocation(srcLat float64,
		srcLng float64) (map[string][]*models.Cab, error)

	GetNearestAvailableCabByType(srcLat float64,
		srcLng float64, cabType string) (*models.Cab, error)

	CreateRide(cab *models.Cab, payload *models.RideRequest,
		eta time.Duration, fareEstimate float64) (*models.Ride, error)

	UpdateRideStatus(rideId string, status string) *models.Ride

	GetDriverEarnings(driverId string) int64
}
