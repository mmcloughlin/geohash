package geohash_test

import (
	"fmt"

	"github.com/mmcloughlin/geohash"
)

func ExampleEncode() {
	fmt.Println(geohash.Encode(48.858, 2.294))
	// Output: u09tunq6qp66
}

func ExampleEncodeInt() {
	fmt.Printf("%016x\n", geohash.EncodeInt(48.858, 2.294))
	// Output: d0139d52c6b54c69
}

func ExampleEncodeIntWithPrecision() {
	fmt.Printf("%08x\n", geohash.EncodeIntWithPrecision(48.858, 2.294, 32))
	// Output: d0139d52
}

func ExampleEncodeWithPrecision() {
	fmt.Println(geohash.EncodeWithPrecision(48.858, 2.294, 5))
	// Output: u09tu
}

func ExampleDecode() {
	lat, lng := geohash.Decode("u09tunq6")
	fmt.Printf("%.3f %.3f\n", lat, lng)
	// Output: 48.858 2.294
}

func ExampleDecodeInt() {
	lat, lng := geohash.DecodeInt(0xd0139d52c6b54c69)
	fmt.Printf("%.3f %.3f\n", lat, lng)
	// Output: 48.858 2.294
}

func ExampleDecodeIntWithPrecision() {
	lat, lng := geohash.DecodeIntWithPrecision(0xd013, uint(16))
	fmt.Printf("%.3f %.3f\n", lat, lng)
	// Output: 48.600 2.000
}
