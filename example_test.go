package geohash_test

import (
	"fmt"

	"github.com/mmcloughlin/geohash"
)

func Example() {
	// Uluru in Australian Outback
	lat, lng := -25.345457, 131.036192

	// Encode a full 12 character string geohash
	fmt.Println(geohash.Encode(lat, lng))

	// Or at lower precision
	fmt.Println(geohash.EncodeWithPrecision(lat, lng, 6))

	// As an integer
	fmt.Printf("%016x\n", geohash.EncodeInt(lat, lng))

	// Decode to a point
	fmt.Println(geohash.Decode("qgmpvf18"))

	// or to a bounding box
	fmt.Println(geohash.BoundingBox("qgmpvf18"))

	// Output:
	// qgmpvf18h86e
	// qgmpvf
	// b3e75db828820cd5
	// -25.3454 131.036
	// {-25.345458984375 -25.345287322998047 131.03599548339844 131.03633880615234}
}

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
