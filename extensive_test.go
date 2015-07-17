package geohash

import (
	"math"
	"testing"
)

// TestCase objects are generated from independent code to verify we get the
// same results. See testcases_test.go.
type TestCase struct {
	hashInt  uint64
	hash     string
	lat, lng float64
}

// Test we get the same string geohashes.
func TestEncode(t *testing.T) {
	for _, c := range testcases {
		hash := Encode(c.lat, c.lng)
		if c.hash != hash {
			t.Errorf("incorrect encode string result for (%v,%v): %s != %s",
				c.lat, c.lng, c.hash, hash)
		}
	}
}

// Test we get the same integer geohashes.
func TestEncodeInt(t *testing.T) {
	for _, c := range testcases {
		hashInt := EncodeInt(c.lat, c.lng)
		if c.hashInt != hashInt {
			t.Errorf("incorrect encode integer result for (%v,%v): %016x != %016x xor %016x",
				c.lat, c.lng, c.hashInt, hashInt, c.hashInt^hashInt)
		}
	}
}

// Verify the prefix property.
func TestPrefixProperty(t *testing.T) {
	for _, c := range testcases {
		for chars := uint(1); chars <= 12; chars++ {
			hash := EncodeWithPrecision(c.lat, c.lng, chars)
			pre := c.hash[:chars]
			if pre != hash {
				t.Errorf("incorrect encode string result for (%v,%v) at precision %d: %s != %s",
					c.lat, c.lng, chars, pre, hash)
			}
		}
	}
}

// Test bounding boxes for string geohashes.
func TestBoundingBox(t *testing.T) {
	for _, c := range testcases {
		box := BoundingBox(c.hash)
		if !box.Contains(c.lat, c.lng) {
			t.Errorf("incorrect bounding box for %s", c.hash)
		}
	}
}

// Test bounding boxes for integer geohashes.
func TestBoundingBoxInt(t *testing.T) {
	for _, c := range testcases {
		box := BoundingBoxInt(c.hashInt)
		if !box.Contains(c.lat, c.lng) {
			t.Errorf("incorrect bounding box for 0x%x", c.hashInt)
		}
	}
}

// Crude test of integer decoding.
func TestDecodeInt(t *testing.T) {
	for _, c := range testcases {
		lat, lng := DecodeInt(c.hashInt)
		if math.Abs(lat-c.lat) > 0.0000001 {
			t.Errorf("large error in decoded latitude for 0x%x", c.hashInt)
		}
		if math.Abs(lng-c.lng) > 0.0000001 {
			t.Errorf("large error in decoded longitude for 0x%x", c.hashInt)
		}
	}
}

type DecodeTestCase struct {
	hash string
	box  Box
}

// Test decoding at various precisions.
func TestDecode(t *testing.T) {
	for _, c := range decodecases {
		lat, lng := Decode(c.hash)
		if !c.box.Contains(lat, lng) {
			t.Errorf("hash %s decoded to %f,%f should lie in %+v",
				c.hash, lat, lng, c.box)
		}
	}
}

// Test roundtrip decoding then encoding again.
func TestDecodeThenEncode(t *testing.T) {
	for _, c := range decodecases {
		precision := uint(len(c.hash))
		lat, lng := Decode(c.hash)
		rehashed := EncodeWithPrecision(lat, lng, precision)
		if c.hash != rehashed {
			t.Errorf("hash %s decoded and re-encoded to %s", c.hash, rehashed)
		}
	}
}
