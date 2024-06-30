package client

func GetRideFareEstimate(point1Lat float64, point1Lng float64,
	point2Lat float64, point2Lng float64) float64 {
	// implement a client to pricing service to get fare estimate.
	dist := GoogleMapClient(point1Lat, point1Lng, point2Lat, point2Lng)

	return dist * 20
}
