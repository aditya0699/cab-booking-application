package client

import (
	"math"
	"time"
)

func GoogleMapClient(point1Lat float64, point1Lng float64,
	point2Lat float64, point2Lng float64) float64 {
	// implement call to google maps api to get distance.
	// using haversine calculator instead.

	const earthRadius = 6371

	lat1Rad := toRadians(point1Lat)
	lng1Rad := toRadians(point1Lng)
	lat2Rad := toRadians(point2Lat)
	lng2Rad := toRadians(point2Lng)

	deltaLat := lat2Rad - lat1Rad
	deltaLng := lng2Rad - lng1Rad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLng/2)*math.Sin(deltaLng/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := earthRadius * c

	return distance

}

func GetETA(point1Lat float64, point1Lng float64,
	point2Lat float64, point2Lng float64, speed float64) time.Duration {
	dist := GoogleMapClient(point1Lat, point1Lng, point2Lat, point2Lng)
	return time.Duration(dist / speed)
}

func toRadians(degree float64) float64 {
	return degree * math.Pi / 180
}
