package main

import (
	"log"
	"net/http"

	"github.com/cab-booking-application/controller"
	"github.com/cab-booking-application/models"
	"github.com/cab-booking-application/repo"
	"github.com/cab-booking-application/service"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	cabs := make(map[string][]*models.Cab)
	rides := make(map[string]*models.Ride)
	repo := repo.NewRepo(cabs, rides)
	srv := service.NewService(repo)
	ctr := controller.NewController(srv)

	router.HandleFunc("/health", ctr.HealthStatus).Methods(http.MethodGet)
	subRouter.HandleFunc("/cabs", ctr.FetchAvailableCabs).Methods(http.MethodGet)
	subRouter.HandleFunc("/ride", ctr.CreateRideRequest).Methods(http.MethodPost)
	subRouter.HandleFunc("/ride/complete/{id}", ctr.CompleteRide).Methods(http.MethodPost)
	subRouter.HandleFunc("/ride/cancel/{id}", ctr.CancelRide).Methods(http.MethodPost)
	subRouter.HandleFunc("driver/earnings/{id}", ctr.DriverEarnings).Methods(http.MethodGet)

	log.Println("Server starting at port 9000")
	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Println("Error in starting server", err)
	}

}
