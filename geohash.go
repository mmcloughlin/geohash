package geohash

import (
	"fmt"
	"math"
)

// Encode the point (lat, lng) as a string geohash with the standard 12
// characters of precision.
func Encode(lat, lng float64) string {
	return EncodeWithPrecision(lat, lng, 12)
}

// Encode the point (lat, lng) as a string geohash with the specified number
// of characters of precision (max 12).
func EncodeWithPrecision(lat, lng float64, chars uint) string {
	bits := 5 * chars
	inthash := EncodeIntWithPrecision(lat, lng, bits)
	enc := base32encoding.Encode(inthash)
	return fmt.Sprintf("%0*s", int(chars), enc)
}

// EncodeInt encodes the point (lat, lng) to a 64-bit integer geohash.
func EncodeInt(lat, lng float64) uint64 {
	latInt := encodeRange(lat, 90)
	lngInt := encodeRange(lng, 180)
	return interleave(latInt, lngInt)
}

// EncodeIntWithPrecision encodes the point (lat, lng) to an integer with the
// specified number of bits.
func EncodeIntWithPrecision(lat, lng float64, bits uint) uint64 {
	hash := EncodeInt(lat, lng)
	return hash >> (64 - bits)
}

type Box struct {
	MinLat float64
	MaxLat float64
	MinLng float64
	MaxLng float64
}

func (b Box) Center() (lat, lng float64) {
	lat = (b.MinLat + b.MaxLat) / 2.0
	lng = (b.MinLng + b.MaxLng) / 2.0
	return
}

func (b Box) Contains(lat, lng float64) bool {
	return (b.MinLat <= lat && lat <= b.MaxLat &&
		b.MinLng <= lng && lng <= b.MaxLng)
}

func errorWithPrecision(bits uint) (latErr, lngErr float64) {
	latBits := bits / 2
	lngBits := bits - latBits
	latErr = 180.0 / math.Exp2(float64(latBits))
	lngErr = 360.0 / math.Exp2(float64(lngBits))
	return
}

func BoundingBox(hash string) Box {
	bits := uint(5 * len(hash))
	inthash := base32encoding.Decode(hash)
	return BoundingBoxIntWithPrecision(inthash, bits)
}

func BoundingBoxIntWithPrecision(hash uint64, bits uint) Box {
	fullHash := hash << (64 - bits)
	latInt, lngInt := deinterleave(fullHash)
	lat := decodeRange(latInt, 90)
	lng := decodeRange(lngInt, 180)
	latErr, lngErr := errorWithPrecision(bits)
	return Box{
		MinLat: lat,
		MaxLat: lat + latErr,
		MinLng: lng,
		MaxLng: lng + lngErr,
	}
}

func Decode(hash string) (lat, lng float64) {
	box := BoundingBox(hash)
	return box.Center()
}

func DecodeIntWithPrecision(hash uint64, bits uint) (lat, lng float64) {
	box := BoundingBoxIntWithPrecision(hash, bits)
	return box.Center()
}

func DecodeInt(hash uint64) (lat, lng float64) {
	return DecodeIntWithPrecision(hash, 64)
}

// Encode the position of x within the range -r to +r as a 32-bit integer.
func encodeRange(x, r float64) uint32 {
	p := (x + r) / (2 * r)
	return uint32(p * math.Exp2(32))
}

// Decode the 32-bit range encoding X back to a value in the range -r to +r.
func decodeRange(X uint32, r float64) float64 {
	p := float64(X) / math.Exp2(32)
	x := 2*r*p - r
	return x
}

// Spread out the 32 bits of x into 64 bits, where the bits of x occupy even
// bit positions.
func spread(x uint32) uint64 {
	X := uint64(x)
	X = (X | (X << 16)) & 0x0000ffff0000ffff
	X = (X | (X << 8)) & 0x00ff00ff00ff00ff
	X = (X | (X << 4)) & 0x0f0f0f0f0f0f0f0f
	X = (X | (X << 2)) & 0x3333333333333333
	X = (X | (X << 1)) & 0x5555555555555555
	return X
}

// Interleave the bits of x and y. In the result, x and y occupy even and odd
// bitlevels, respectively.
func interleave(x, y uint32) uint64 {
	return spread(x) | (spread(y) << 1)
}

// Squash the even bitlevels of X into a 32-bit word. Odd bitlevels of X are
// ignored, and may take any value.
func squash(X uint64) uint32 {
	X &= 0x5555555555555555
	X = (X | (X >> 1)) & 0x3333333333333333
	X = (X | (X >> 2)) & 0x0f0f0f0f0f0f0f0f
	X = (X | (X >> 4)) & 0x00ff00ff00ff00ff
	X = (X | (X >> 8)) & 0x0000ffff0000ffff
	X = (X | (X >> 16)) & 0x00000000ffffffff
	return uint32(X)
}

// Deinterleave the bits of X into 32-bit words containing the even and odd
// bitlevels of X, respectively.
func deinterleave(X uint64) (uint32, uint32) {
	return squash(X), squash(X >> 1)
}
