package service

import (
	"log"

	"github.com/cab-booking-application/client"
	"github.com/cab-booking-application/constants"
	"github.com/cab-booking-application/models"
	"github.com/cab-booking-application/repo"
)

type BookingService struct {
	repo repo.IBookingRepo
}

func NewService(repo repo.IBookingRepo) IBookingService {
	return &BookingService{
		repo: repo,
	}
}

func (srv *BookingService) FetchAvailableCabs(srcLat float64,
	srcLng float64) (map[string][]*models.Cab, error) {

	return srv.repo.GetAllCabsByLocation(srcLat, srcLng)
}

func (srv *BookingService) CreateRide(payload *models.RideRequest) (*models.Ride, error) {
	cab, err := srv.repo.GetNearestAvailableCabByType(payload.SrcLat, payload.SrcLng,
		payload.CabType)
	if err != nil {
		log.Println("Error in finding cab")
		return nil, err
	}
	eta := client.GetETA(payload.SrcLat, payload.SrcLng,
		cab.Location.Lat, cab.Location.Lng, cab.Speed)
	fareEstimate := client.GetRideFareEstimate(payload.SrcLat, payload.SrcLng,
		cab.Location.Lat, cab.Location.Lng)
	ride, err := srv.repo.CreateRide(cab, payload, eta, fareEstimate)
	if err != nil {
		log.Println("Error in creating ride")
		return nil, err
	}
	return ride, err
}

func (srv *BookingService) CompleteRide(rideID string) *models.Ride {

	return srv.repo.UpdateRideStatus(rideID, constants.Completed)
}

func (srv *BookingService) CancelRide(rideID string) *models.Ride {

	return srv.repo.UpdateRideStatus(rideID, constants.Cancelled)
}

func (srv *BookingService) GetDriverEarnings(driverId string) int64 {
	return srv.repo.GetDriverEarnings(driverId)
}

func (srv *BookingService) GetRideInvoice(rideId string) *models.PricingTier {
	return srv.repo.GetRideInvoice(rideId)
}
