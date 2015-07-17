package geohash

import "fmt"

func ExampleEncode() {
	fmt.Println(Encode(48.858, 2.294))
	// Output: u09tunq6qp66
}
