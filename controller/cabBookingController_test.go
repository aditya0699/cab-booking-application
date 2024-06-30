package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cab-booking-application/constants"
	"github.com/cab-booking-application/controller"
	"github.com/cab-booking-application/models"
	service_mocks "github.com/cab-booking-application/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHealthStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mocks.NewMockIBookingService(ctrl)
	bc := controller.NewController(mockService)

	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(bc.HealthStatus)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response["status"])
}

func TestFetchAvailableCabs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mocks.NewMockIBookingService(ctrl)
	bc := controller.NewController(mockService)

	// Assuming the FetchAvailableCabs function in the service returns an array of cabs
	mockService.EXPECT().FetchAvailableCabs(37.7749, -122.4194).Return([]models.Cab{}, nil)

	req, err := http.NewRequest("GET", "/cabs?srclat=37.7749&srclng=-122.4194", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(bc.FetchAvailableCabs)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateRideRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mocks.NewMockIBookingService(ctrl)
	bc := controller.NewController(mockService)

	rideRequest := &models.RideRequest{
		SrcLat:  37.7749,
		SrcLng:  -122.4194,
		DestLat: 37.7849,
		DestLng: -122.4094,
		CabType: constants.Sedan,
	}

	mockService.EXPECT().CreateRide(rideRequest).Return(&models.Ride{}, nil)

	payload, err := json.Marshal(rideRequest)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/ride", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(bc.CreateRideRequest)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
