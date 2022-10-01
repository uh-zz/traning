package distance

import "math"

// Coordinate 緯度経度
type Coordinate struct {
	Longitude float64
	Latitude  float64
}

// EarthRadius 赤道半径
const EarthRadius = 6378140

// DistanceOnTheEarth 地球上の2点間の距離を出す（球面三角法）
func DistanceOnTheEarth(from, to Coordinate) float64 {

	fromLadLon := from.Longitude * math.Pi / 180
	fromLadLat := from.Latitude * math.Pi / 180

	toLadLon := to.Longitude * math.Pi / 180
	toLadLat := to.Latitude * math.Pi / 180

	alpha := math.Sin(fromLadLat)*math.Sin(toLadLat) +
		math.Cos(fromLadLat)*math.Cos(toLadLat)*math.Cos(fromLadLon-toLadLon)

	arcAlpha := math.Acos(alpha)

	return arcAlpha * EarthRadius / 1000
}
