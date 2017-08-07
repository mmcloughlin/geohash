package geohash

import "testing"

func randomPoints(n int) [][2]float64 {
	var points [][2]float64
	for i := 0; i < n; i++ {
		lat, lon := RandomPoint()
		points = append(points, [2]float64{lat, lon})
	}
	return points
}

func BenchmarkEncode(b *testing.B) {
	points := randomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(points[i][0], points[i][1])
	}
}

func BenchmarkEncodeInt(b *testing.B) {
	points := randomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeInt(points[i][0], points[i][1])
	}
}
