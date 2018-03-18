package geohash

//go:noescape
func EncodeInt(lat, lng float64) uint64
