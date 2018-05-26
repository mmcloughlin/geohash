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

func RandomPoints(n int) [][2]float64 {
	points := make([][2]float64, n)
	for i := 0; i < n; i++ {
		lat, lng := RandomPoint()
		points[i] = [2]float64{lat, lng}
	}
	return points
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

func RandomStringGeohashWithPrecision(chars uint) string {
	const alphabet = "0123456789bcdefghjkmnpqrstuvwxyz"
	b := make([]byte, chars)
	for i := uint(0); i < chars; i++ {
		b[i] = alphabet[rand.Intn(32)]
	}
	return string(b)
}

func RandomStringGeohashesWithPrecision(n int, chars uint) []string {
	geohashes := make([]string, n)
	for i := 0; i < n; i++ {
		geohashes[i] = RandomStringGeohashWithPrecision(chars)
	}
	return geohashes
}

func RandomIntGeohash() uint64 {
	return (uint64(rand.Uint32()) << 32) | uint64(rand.Uint32())
}

func RandomIntGeohashes(n int) []uint64 {
	geohashes := make([]uint64, n)
	for i := 0; i < n; i++ {
		geohashes[i] = RandomIntGeohash()
	}
	return geohashes
}
