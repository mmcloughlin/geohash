package geohash_test

import (
	"fmt"

	"github.com/mmcloughlin/geohash"
)

func ExampleEncode() {
	fmt.Println(geohash.Encode(48.858, 2.294))
	// Output: u09tunq6qp66
}
