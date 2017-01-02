# geohash

Go [geohash](https://en.wikipedia.org/wiki/Geohash) library offering encoding
and decoding for string and integer geohashes.

[![GoDoc Reference](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](http://godoc.org/github.com/mmcloughlin/geohash)
[![Build status](https://img.shields.io/travis/mmcloughlin/geohash.svg?style=flat-square)](https://travis-ci.org/mmcloughlin/geohash)
[![Coverage](https://img.shields.io/coveralls/mmcloughlin/geohash.svg?style=flat-square)](https://coveralls.io/r/mmcloughlin/geohash)
[![Go Report Card](https://goreportcard.com/badge/github.com/mmcloughlin/geohash?style=flat-square)](https://goreportcard.com/report/github.com/mmcloughlin/geohash)

## Install

Fetch the package with

```
go get github.com/mmcloughlin/geohash
```

And import it into your programs with

```go
import github.com/mmcloughlin/geohash
```

## Usage

#### func  Decode

```go
func Decode(hash string) (lat, lng float64)
```
Decode the string geohash to a (lat, lng) point.

#### func  DecodeInt

```go
func DecodeInt(hash uint64) (lat, lng float64)
```
DecodeInt decodes the provided 64-bit integer geohash to a (lat, lng) point.

#### func  DecodeIntWithPrecision

```go
func DecodeIntWithPrecision(hash uint64, bits uint) (lat, lng float64)
```
DecodeIntWithPrecision decodes the provided integer geohash with bits of
precision to a (lat, lng) point.

#### func  Encode

```go
func Encode(lat, lng float64) string
```
Encode the point (lat, lng) as a string geohash with the standard 12 characters
of precision.

#### func  EncodeInt

```go
func EncodeInt(lat, lng float64) uint64
```
EncodeInt encodes the point (lat, lng) to a 64-bit integer geohash.

#### func  EncodeIntWithPrecision

```go
func EncodeIntWithPrecision(lat, lng float64, bits uint) uint64
```
EncodeIntWithPrecision encodes the point (lat, lng) to an integer with the
specified number of bits.

#### func  EncodeWithPrecision

```go
func EncodeWithPrecision(lat, lng float64, chars uint) string
```
EncodeWithPrecision encodes the point (lat, lng) as a string geohash with the
specified number of characters of precision (max 12).

#### type Box

```go
type Box struct {
	MinLat float64
	MaxLat float64
	MinLng float64
	MaxLng float64
}
```

Box represents a rectangle in latitude/longitude space.

#### func  BoundingBox

```go
func BoundingBox(hash string) Box
```
BoundingBox returns the region encoded by the given string geohash.

#### func  BoundingBoxInt

```go
func BoundingBoxInt(hash uint64) Box
```
BoundingBoxInt returns the region encoded by the given 64-bit integer geohash.

#### func  BoundingBoxIntWithPrecision

```go
func BoundingBoxIntWithPrecision(hash uint64, bits uint) Box
```
BoundingBoxIntWithPrecision returns the region encoded by the integer geohash
with the specified precision.

#### func (Box) Center

```go
func (b Box) Center() (lat, lng float64)
```
Center returns the center of the box.

#### func (Box) Contains

```go
func (b Box) Contains(lat, lng float64) bool
```
Contains decides whether (lat, lng) is contained in the box. The containment
test is inclusive of the edges and corners.

#### func (Box) Round

```go
func (b Box) Round() (lat, lng float64)
```
Round returns a point inside the box, making an effort to round to minimal
precision.

## Alternatives

* https://github.com/gansidui/geohash
* https://github.com/TomiHiltunen/geohash-golang
* https://github.com/corsc/go-geohash
* https://github.com/pierrre/geohash
* https://github.com/keshavdv/go-geohash
 
## Reference

* https://github.com/yinqiwen/geohash-int
* https://github.com/hkwi/python-geohash
* https://github.com/sunng87/node-geohash
* https://github.com/davetroy/geohash-js
