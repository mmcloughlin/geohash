package geohash

import "testing"

func TestInterleave(t *testing.T) {
	res := interleave(0xf80, 0xbc9)
	if 0xdfe082 != res {
		t.Errorf("incorrect interleave result")
	}
}
