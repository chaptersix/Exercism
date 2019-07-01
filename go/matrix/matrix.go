package matrix

import (
	"errors"
	"strconv"
	"strings"
)

//Matrix is a 2d array of ints
type Matrix [][]int

// New create a 2d int array for the provided matrix string
func New(matrixString string) (Matrix, error) {
	rows := strings.Split(matrixString, "\n")
	m := make([][]int, len(rows))
	for i, row := range rows {
		columnStrings := strings.Split(strings.TrimSpace(row), " ")
		rowInts := make([]int, len(columnStrings))
		for j, cell := range columnStrings {
			cellInt, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			rowInts[j] = cellInt
		}
		m[i] = rowInts
		if i > 0 && len(m[i-1]) != len(m[i]) {
			return nil, errors.New("rows of different lengths")
		}
	}
	return m, nil
}

//Rows returns  the given Matrix
func (m Matrix) Rows() [][]int {
	new := make([][]int, len(m))
	for index, row := range m {
		new[index] = make([]int, len(row))
		copy(new[index], row)
	}
	return new
}

//Cols returns  the given Matrix rotated
func (m Matrix) Cols() [][]int {
	col := make([][]int, len(m[0]))
	mRow := len(m[0])
	mCol := len(m)
	for i := 0; i < mRow; i++ {
		row := make([]int, mCol)
		for j := 0; j < mCol; j++ {
			row[j] = m[j][i]
		}
		col[i] = row
	}
	return col
}

//Set changes the cell to value at the given coordinates
func (m Matrix) Set(row, col, value int) bool {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[0]) {
		return false
	}
	m[row][col] = value
	return true
}
