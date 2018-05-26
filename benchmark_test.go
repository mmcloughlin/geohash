// +build go1.7

package geohash

import (
	"strconv"
	"testing"
)

func BenchmarkEncodeInt(b *testing.B) {
	points := RandomPoints(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeInt(points[i%1024][0], points[i%1024][1])
	}
}

func BenchmarkEncode(b *testing.B) {
	points := RandomPoints(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(points[i%1024][0], points[i%1024][1])
	}
}

func BenchmarkEncodeWithPrecision(b *testing.B) {
	points := RandomPoints(1024)
	for chars := uint(1); chars <= 12; chars++ {
		name := strconv.FormatUint(uint64(chars), 10)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EncodeWithPrecision(points[i%1024][0], points[i%1024][1], chars)
			}
		})
	}
}

func BenchmarkDecodeInt(b *testing.B) {
	geohashes := RandomIntGeohashes(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DecodeInt(geohashes[i%1024])
	}
}

func BenchmarkDecode(b *testing.B) {
	for chars := uint(1); chars <= 12; chars++ {
		name := strconv.FormatUint(uint64(chars), 10)
		b.Run(name, func(b *testing.B) {
			geohashes := RandomStringGeohashesWithPrecision(1024, chars)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Decode(geohashes[i%1024])
			}
		})
	}
}
