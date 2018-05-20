package geohash

var EncodeInt = encodeIntGeneric

func init() {
	EncodeInt = encodeIntAsm
}

//go:noescape
func encodeIntAsm(lat, lng float64) uint64
