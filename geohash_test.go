package geohash

import (
	"math"
	"math/rand"
	"testing"
	"testing/quick"
)

func TestInterleaving(t *testing.T) {
	cases := []struct {
		x, y uint32
		X    uint64
	}{
		{0x00000000, 0x00000000, 0x0000000000000000},
		{0xffffffff, 0x00000000, 0x5555555555555555},
		{0x789e22e9, 0x8ed4182e, 0x95e8e37406845ce9},
		{0xb96346bb, 0xf8a80f02, 0xefc19c8510be454d},
		{0xa1dfc6c2, 0x01c886f9, 0x4403f1d5d03cfa86},
		{0xfb59e296, 0xad2c6c02, 0xdde719e17ca4411c},
		{0x94e0bbf2, 0xb520e8b2, 0xcb325c00edc5df0c},
		{0x1638ca5f, 0x5e16a514, 0x23bc0768d8661375},
		{0xe15bbbf7, 0x0f6bf376, 0x54ab39cfef4f7f3d},
		{0x06a476a7, 0x94f35ec7, 0x8234ee1a37bce43f},
	}
	for _, c := range cases {
		res := interleave(c.x, c.y)
		if c.X != res {
			t.Errorf("incorrect interleave result")
		}

		x, y := deinterleave(c.X)
		if c.x != x || c.y != y {
			t.Errorf("incorrect deinterleave result")
		}
	}
}

func TestBase32Decode(t *testing.T) {
	x := base32encoding.Decode("ezs42")
	if 0xdfe082 != x {
		t.Errorf("incorrect base64 decoding")
	}
}

func TestBase32Encode(t *testing.T) {
	s := base32encoding.Encode(0xdfe082)
	if "ezs42" != s {
		t.Errorf("incorrect base64 encoding")
	}
}

func RandomBox() Box {
	lat1 := -90 + 180*rand.Float64()
	lat2 := -90 + 180*rand.Float64()
	lng1 := -180 + 360*rand.Float64()
	lng2 := -180 + 360*rand.Float64()
	return Box{
		MinLat: math.Min(lat1, lat2),
		MaxLat: math.Max(lat1, lat2),
		MinLng: math.Min(lng1, lng2),
		MaxLng: math.Max(lng1, lng2),
	}
}

func TestBoxRound(t *testing.T) {
	f := func() bool {
		b := RandomBox()
		lat, lng := b.Round()
		return b.Contains(lat, lng)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBoxCenter(t *testing.T) {
	b := Box{
		MinLat: 1,
		MaxLat: 2,
		MinLng: 3,
		MaxLng: 4,
	}
	lat, lng := b.Center()
	if 1.5 != lat || 3.5 != lng {
		t.Errorf("incorrect box center")
	}
}

func TestBoxContains(t *testing.T) {
	b := Box{
		MinLat: 1,
		MaxLat: 2,
		MinLng: 3,
		MaxLng: 4,
	}
	cases := []struct {
		lat, lng float64
		expect   bool
	}{
		{1.5, 3.5, true},
		{0.5, 3.5, false},
		{7.0, 3.5, false},
		{1.5, 1.5, false},
		{1.5, 9.5, false},
		{1, 3, true},
		{1, 4, true},
		{2, 3, true},
		{2, 4, true},
	}
	for _, c := range cases {
		if c.expect != b.Contains(c.lat, c.lng) {
			t.Errorf("contains %f,%f should be %t", c.lat, c.lng, c.expect)
		}
	}
}

func TestWikipediaExample(t *testing.T) {
	h := EncodeWithPrecision(42.6, -5.6, 5)
	if "ezs42" != h {
		t.Errorf("incorrect encoding")
	}
}

func TestLeadingZero(t *testing.T) {
	h := EncodeWithPrecision(-74.761330, -140.309714, 6)
	if 6 != len(h) {
		t.Errorf("incorrect geohash length")
	}
	if "0fsnxn" != h {
		t.Errorf("incorrect encoding")
	}
}
