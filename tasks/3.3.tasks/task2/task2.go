package task2

import "math"

func Distance(lat1, lon1 float32, lat2, lon2 float32) float32 {
	R := 6371.0

	lat1Rad := lat1 * (math.Pi / 180.0)
	lon1Rad := lon1 * (math.Pi / 180.0)
	lat2Rad := lat2 * (math.Pi / 180.0)
	lon2Rad := lon2 * (math.Pi / 180.0)

	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	a := math.Sin(float64(dlat/2))*math.Sin(float64(dlat/2)) + math.Cos(float64(lat1Rad))*math.Cos(float64(lat2Rad))*math.Sin(float64(dlon/2))*math.Sin(float64(dlon/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c

	return float32(distance)
}
