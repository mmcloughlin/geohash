package geohash

import (
	"testing"
)

const lat, lng = 33.0, -72.0

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Encode(lat, lng)
	}
}

func BenchmarkEncodeInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = EncodeInt(lat, lng)
	}
}
