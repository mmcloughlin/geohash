// +build !amd64

package geohash

// EncodeInt encodes the point (lat, lng) to a 64-bit integer geohash.
func EncodeInt(lat, lng float64) uint64 {
	return encodeIntGeneric(lat, lng)
}
