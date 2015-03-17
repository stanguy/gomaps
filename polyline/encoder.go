package polyline

// Utility for creating Google Maps Encoded GPolylines
//
// License: You may distribute this code under the same terms as Ruby itself
//
// Author: Joel Rosenberg
//
// ( Drawing from the official example pages as well as Mark McClure's work )

import (
	"math"
)

func Encode(points [][]float64) string {

	plat := int32(0)
	plon := int32(0)

	encoded := ""

	for _, p := range points {
		late5 := int32(math.Floor(p[0] * 1e5))
		lone5 := int32(math.Floor(p[1] * 1e5))

		dlat := late5 - plat
		dlon := lone5 - plon

		plat = late5
		plon = lone5

		encoded += polyline_encode_signed_number(dlat)
		encoded += polyline_encode_signed_number(dlon)

	}
	return encoded
}

func polyline_encode_signed_number(num int32) string {
	sgn_num := num << 1
	if num < 0 {
		sgn_num = ^sgn_num
	}
	return polyline_encode_number(sgn_num)
}

func polyline_encode_number(num int32) string {

	encoded := ""

	for num >= 0x20 {
		encoded += string(rune(((num & 0x1f) | 0x20) + 63))
		num = num >> 5
	}
	encoded += string(rune(num + 63))

	return encoded
}
