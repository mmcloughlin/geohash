package geohash

import "testing"

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
