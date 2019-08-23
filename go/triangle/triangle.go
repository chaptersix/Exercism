package triangle

import (
	"math"
	"sort"
)
//Kind is an enum of the types of triangles
type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) (k Kind) {
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	if math.IsNaN(sides[0]) || math.IsInf(sides[0], -1) || math.IsInf(sides[2], 1) {
		return NaT
	}
	if sides[0] <= 0 || sides[0]+sides[1] < sides[2] {
		k = NaT
		return
	}
	if sides[0] == sides[2] {
		k = Equ
		return
	}
	if sides[0] == sides[1] || sides[1] == sides[2] {
		k = Iso
		return
	}
	k = Sca
	return
}
