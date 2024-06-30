# Cab Booking Application

## Overview

This is a simple cab booking application implemented in Go. The application provides APIs for booking cabs, fetching available cabs, completing and canceling rides, and checking driver earnings.

## Prerequisites

- Go 1.16 or later
- Gorilla Mux

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/your-repo/cab-booking-application.git
    cd cab-booking-application
    ```

2. Install dependencies:
    ```sh
    go get -u github.com/gorilla/mux
    ```

## Usage

1. Build and run the application:
    ```sh
    go run main.go
    ```

2. The server will start at port 9000. You can access the APIs using the following endpoints:

    - Health Check:
        ```sh
        GET /health
        ```
        Returns the health status of the application.

    - Fetch Available Cabs:
        ```sh
        GET /api/v1/cabs?srclat={latitude}&srclng={longitude}
        ```
        Fetches available cabs near the provided source latitude and longitude.

    - Create Ride Request:
        ```sh
        POST /api/v1/ride
        ```
        Creates a new ride request. The request body should contain the ride details in JSON format.

    - Complete Ride:
        ```sh
        POST /api/v1/ride/complete/{id}
        ```
        Completes a ride with the given ride ID.

    - Cancel Ride:
        ```sh
        POST /api/v1/ride/cancel/{id}
        ```
        Cancels a ride with the given ride ID.

    - Driver Earnings:
        ```sh
        GET /api/v1/driver/earnings/{id}
        ```
        Fetches the earnings for the driver with the given driver ID.

## Project Structure

- `main.go`: Entry point of the application.
- `controller/`: Contains the HTTP handlers for the APIs.
- `models/`: Contains the data models for the application.
- `repo/`: Contains the repository layer for data persistence.
- `service/`: Contains the business logic for the application.

## Running Tests

1. Install `mockgen`:
    ```sh
    go install github.com/golang/mock/mockgen@v1.6.0
    ```

2. Generate mocks:
    ```sh
    mockgen -source=service/booking_service.go -destination=service/mock_booking_service.go -package=service
    ```

3. Run tests:
    ```sh
    go test ./...
    ```

Note: Used in memory datastores for ease, we can definitely use some persistent datastores.
