package geohash

import (
	"math"
	"math/rand"
)

func RandomPoint() (lat, lng float64) {
	lat = -90 + 180*rand.Float64()
	lng = -180 + 360*rand.Float64()
	return
}

func RandomBox() Box {
	lat1, lng1 := RandomPoint()
	lat2, lng2 := RandomPoint()
	return Box{
		MinLat: math.Min(lat1, lat2),
		MaxLat: math.Max(lat1, lat2),
		MinLng: math.Min(lng1, lng2),
		MaxLng: math.Max(lng1, lng2),
	}
}
