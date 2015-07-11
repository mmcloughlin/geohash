package geohash

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
