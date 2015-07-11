package geohash

import "testing"

func TestInterleave(t *testing.T) {
	res := interleave(0xf80, 0xbc9)
	if 0xdfe082 != res {
		t.Errorf("incorrect interleave result")
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
