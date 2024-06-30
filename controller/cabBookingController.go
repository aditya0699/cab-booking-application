package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/cab-booking-application/models"
	"github.com/cab-booking-application/service"
	"github.com/gorilla/mux"
)

type BookingController struct {
	bookingService service.IBookingService
}

func NewController(bookingService service.IBookingService) *BookingController {
	return &BookingController{
		bookingService: bookingService,
	}
}

func (ctr *BookingController) HealthStatus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
	w.WriteHeader(http.StatusOK)
}

func (ctr *BookingController) FetchAvailableCabs(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	srcLat := queryParams.Get("srclat")
	srcLng := queryParams.Get("srclng")
	if srcLat == "" || srcLng == "" {
		log.Println("Missing required query params")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	srcLatFloat, err := strconv.ParseFloat(srcLat, 64)
	if err != nil {
		log.Println("Error in parsing srcLat")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	srcLngFloat, err := strconv.ParseFloat(srcLng, 64)
	if err != nil {
		log.Println("Error in parsing srcLng")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	availableCabs, err := ctr.bookingService.FetchAvailableCabs(srcLatFloat, srcLngFloat)
	if err != nil {
		log.Println("Error in getting cabs")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(availableCabs)
}

func (ctr *BookingController) CreateRideRequest(w http.ResponseWriter, r *http.Request) {
	payloadBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in parsing request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var payload *models.RideRequest
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		log.Println("Error in parsing payload")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ride, err := ctr.bookingService.CreateRide(payload)
	if err != nil {
		log.Println("Error in creating ride")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ride)
}

func (ctr *BookingController) CompleteRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rideId, ok := params["id"]
	if !ok {
		log.Println("Error in parsing rideId")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ride := ctr.bookingService.CompleteRide(rideId)
	if ride == nil {
		log.Println("Error in completing ride")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "success", "status": "Ok", "ride": *ride})
}

func (ctr *BookingController) CancelRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rideId, ok := params["id"]
	if !ok {
		log.Println("Error in parsing rideId")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ride := ctr.bookingService.CancelRide(rideId)
	if ride != nil {
		log.Println("Error in completing ride")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "success", "status": "Ok", "ride": *ride})
}

func (ctr *BookingController) DriverEarnings(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	driverId, ok := params["id"]
	if !ok {
		log.Println("Error in parsing driverId")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	earnings := ctr.bookingService.GetDriverEarnings(driverId)
	json.NewEncoder(w).Encode(earnings)
}

func (ctr *BookingController) GetRideInvoice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rideId, ok := params["id"]
	if !ok {
		log.Println("Error in parsing driverId")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rideFare := ctr.bookingService.GetRideInvoice(rideId)
	json.NewEncoder(w).Encode(rideFare)
}
