package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 緯度経度
	fromLocation := GeoLocation("東京")
	toLocation := GeoLocation("大阪")

	fromLon, _ := strconv.ParseFloat(fromLocation[0].Longitude, 64)
	fromLat, _ := strconv.ParseFloat(fromLocation[0].Latitude, 64)

	toLon, _ := strconv.ParseFloat(toLocation[0].Longitude, 64)
	toLat, _ := strconv.ParseFloat(toLocation[0].Latitude, 64)

	from := Coordinate{
		Longitude: fromLon,
		Latitude:  fromLat,
	}

	to := Coordinate{
		Longitude: toLon,
		Latitude:  toLat,
	}

	// 2点間の距離計算
	res := DistanceOnTheEarth(from, to)

	fmt.Println("from", fromLocation[0].City, fromLocation[0].Town)
	fmt.Println("to", toLocation[0].City, toLocation[0].Town)
	fmt.Println(res, "km")
}
