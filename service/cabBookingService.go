package service

import "github.com/cab-booking-application/models"

type IBookingService interface {
	FetchAvailableCabs(srcLat float64,
		srcLng float64) (map[string][]*models.Cab, error)

	CreateRide(payload *models.RideRequest) (*models.Ride, error)

	CompleteRide(rideID string) *models.Ride

	CancelRide(rideID string) *models.Ride

	GetDriverEarnings(driverId string) int64
}
