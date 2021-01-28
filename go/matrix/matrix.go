package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix contains a matrix.
type Matrix [][]int

// New generates a matrix with the given input.
// Error is thrown if invalid input or if a given
// number overflows int.
func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	colLen := -1
	ret := make(Matrix, len(rows))
	for i, r := range rows {
		cols := strings.Split(strings.TrimSpace(r), " ")
		ln := len(cols)
		if colLen == -1 {
			colLen = ln
		}
		if ln != colLen || ln == 0 {
			return nil, errors.New("Invalid row")
		}
		ret[i] = make([]int, colLen)
		for j, c := range cols {
			val, err := strconv.Atoi(c)
			if err != nil {
				return nil, errors.New("Invalid integer")
			}
			ret[i][j] = val
		}
	}
	return ret, nil
}

// Cols does a rotation and returns a copy
func (m Matrix) Cols() [][]int {
	ret := make(Matrix, len(m[0]))
	for _, r := range m {
		for j, c := range r {
			ret[j] = append(ret[j], c)
		}
	}
	return ret
}

// Rows returns a copy of the matrix
func (m Matrix) Rows() [][]int {
	ret := make(Matrix, len(m))
	for i, r := range m {
		ret[i] = make([]int, len(r))
		for j, c := range r {
			ret[i][j] = c
		}
	}
	return ret
}

// Set changes an element at given coords
func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 {
		return false
	}
	if r >= len(*m) || c >= len((*m)[0]) {
		return false
	}
	(*m)[r][c] = val
	return true
}
