package grains

import (
	"errors"
	"math"
)

// Square calculates the number of grains on a given tile
func Square(tileNumber int) (uint64, error) {
	if tileNumber <= 0 || tileNumber > 64 {
		return 0, errors.New("tileNumber is not between 1 and 64")
	}
	grains := uint64(math.Pow(2, float64(tileNumber-1)))
	return grains, nil
}

// Total calculates the total number of grains on a chessboard
func Total() uint64 {
	// this does not require a calculation based on input. total number is  costaint
	return 18446744073709551615
}
