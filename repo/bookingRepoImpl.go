package repo

import (
	"fmt"
	"time"

	"github.com/cab-booking-application/client"
	"github.com/cab-booking-application/constants"
	"github.com/cab-booking-application/models"
	uuid "github.com/satori/go.uuid"
)

type BookingRepo struct {
	Cabs    map[string][]*models.Cab
	Rides   map[string]*models.Ride
	Drivers map[string]*models.Driver
	Users   map[string]*models.User
}

func NewRepo(cabs map[string][]*models.Cab, rides map[string]*models.Ride) IBookingRepo {
	return &BookingRepo{
		Cabs:  cabs,
		Rides: rides,
	}
}

func (repo *BookingRepo) GetAllCabsByLocation(srcLat float64,
	srcLng float64) (map[string][]*models.Cab, error) {

	availableCabsByTier := make(map[string][]*models.Cab)
	for tier, Cabs := range repo.Cabs {
		cabs := make([]*models.Cab, 0)
		for _, cab := range Cabs {
			if client.GoogleMapClient(srcLat, srcLng,
				cab.Location.Lat, cab.Location.Lng) <= constants.Radius {
				cabs = append(cabs, cab)
			}
		}
		availableCabsByTier[tier] = cabs
	}
	return availableCabsByTier, nil
}

func (repo *BookingRepo) GetNearestAvailableCabByType(srcLat float64,
	srcLng float64, cabType string) (*models.Cab, error) {

	for _, cab := range repo.Cabs[cabType] {
		if cab.Driver.Status == constants.StatusAvailable && client.GoogleMapClient(srcLat, srcLng,
			cab.Location.Lat, cab.Location.Lng) <= constants.Radius {
			return cab, nil
		}
	}
	return nil, fmt.Errorf("No Cab Available at the time")
}

func (repo *BookingRepo) CreateRide(cab *models.Cab,
	payload *models.RideRequest, eta time.Duration, fareEstimate float64) (*models.Ride, error) {

	ride := &models.Ride{
		ID: uuid.NewV4().String(),
		// get userId from jwt token, using uuid here.
		UserID:   uuid.NewV4().String(),
		DriverID: cab.Driver.ID,
		Pickup: &models.Location{
			Lat: payload.SrcLat,
			Lng: payload.SrcLng,
		},
		Drop: &models.Location{
			Lat: payload.DestLat,
			Lng: payload.DestLng,
		},
		CabType:     cab.Type,
		Status:      constants.Accepted,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(eta),
		TotalAmount: fareEstimate,
	}
	repo.Rides[ride.ID] = ride

	for _, Cabs := range repo.Cabs[payload.CabType] {
		if Cabs.ID == cab.ID {
			cab.Driver.Status = constants.StatusInRide
		}
	}
	return ride, nil

}

func (repo *BookingRepo) UpdateRideStatus(rideId string, status string) *models.Ride {
	ride := repo.Rides[rideId]
	ride.Status = status

	if status == constants.Completed {
		driver := repo.Drivers[ride.DriverID]
		driver.TotalEarnings += ride.TotalAmount
		for _, Cabs := range repo.Cabs[ride.CabType] {
			if Cabs.ID == ride.CabId {
				Cabs.Driver.Status = constants.StatusAvailable
				break
			}
		}
	}

	if status == constants.Cancelled && ride.StartTime.UnixMilli() <
		(time.Now().UnixMilli()-(int64)(5*60*1000)) {
		ride.TotalAmount = constants.CancellationFine

		for _, Cabs := range repo.Cabs[ride.CabType] {
			if Cabs.ID == ride.CabId {
				Cabs.Driver.Status = constants.StatusAvailable
				break
			}
		}
	}
	repo.Rides[rideId] = ride

	return ride
}

func (repo *BookingRepo) GetDriverEarnings(driverId string) int64 {
	return int64(repo.Drivers[driverId].TotalEarnings)
}

func (repo *BookingRepo) GetRideInvoice(rideId string) *models.PricingTier {
	ride := repo.Rides[rideId]
	return &models.PricingTier{
		PerKM:    10,
		BaseFare: 50,
		Tier:     ride.CabType,
	}
}
