package minesweeper

import (
	"errors"
	"fmt"
)

func isValidHeader(row []byte) bool {
	// must start and end with a +
	if row[0] != '+' || row[len(row)-1] != '+' {
		return false
	}

	// checking that others chars are '-'
	l := len(row)
	if l < 3 {
		return true
	}

	for i := 1; i < l-2; i++ {
		if row[i] != '-' {
			return false
		}
	}

	return true
}

func isValidRow(row []byte) bool {
	l := len(row)
	if row[0] != '|' || row[l-1] != '|' {
		return false
	}

	if l < 3 {
		return true
	}

	// check others chars
	for i := 1; i < l-2; i++ {
		if row[i] != '*' && row[i] != ' ' {
			return false
		}
	}

	return true
}

func (b Board) isValid() bool {
	// if no rows, error
	l := len(b)
	fmt.Printf("len(b): %d\n", l)
	if l == 0 {
		return false
	}

	rowsLen := 0
	for rowIndex, row := range b {
		rowLen := len(row)
		// first row
		if rowIndex == 0 || rowIndex == l-1 {
			if !isValidHeader(row) {
				fmt.Println("invalid header/footer")
				return false
			}
		} else {
			if !isValidRow(row) {
				fmt.Println("invalid row")
				return false
			}
		}

		// rowsLen check
		if rowsLen == 0 {
			rowsLen = rowLen
			continue
		}
		if rowLen != rowsLen || rowLen == 0 {
			return false
		}

	}
	return true
}

// returns only the inner matrix we want to work with
func (b Board) getMatrix() (res [][]byte, err error) {
	if !b.isValid() {
		err = errors.New("Invalid input")
		return
	}
	ySize := len(b)
	fmt.Printf("ySize: %d\n", ySize)
	xSize := len(b[0])
	fmt.Printf("xSize: %d\n", xSize)

	if ySize == 2 || xSize == 2 {
		fmt.Println("nothing to compute")
		return
	}

	for y := 1; y < ySize-1; y++ {
		var row []byte
		for x := 1; x < xSize-1; x++ {
			row = append(row, b[y][x])
		}
		res = append(res, row)
	}

	return
}

// Count does a count
func (b *Board) Count() error {
	matrix, err := b.getMatrix()
	if err != nil {
		return err
	}

	fmt.Printf("Matrix: \n%v\n", matrix)

	// for each matrix line, compute each cell mines in the nearing
	ySize := len(matrix)
	for y, row := range matrix {
		isFirstLine := y == 0
		isLastLine := y == ySize-1
		for x, cell := range row {
			isFirstCell := x == 0
			isLastCell := x == len(row)-1
			if cell == '*' {
				fmt.Println("Bomb here")
				continue
			}
			// ceil should contain a ' ', then we compute the replacing number
			bombCount := 0
			// first line
			if !isFirstLine {
				// check the line before us
				if !isFirstCell {
					if matrix[y-1][x-1] == '*' {
						bombCount++
					}
				}
				if matrix[y-1][x] == '*' {
					bombCount++
				}
				if !isLastCell {
					if matrix[y-1][x+1] == '*' {
						bombCount++
					}
				}
			}
			// current line
			if !isFirstCell {
				if matrix[y][x-1] == '*' {
					bombCount++
				}
			}
			if !isLastCell {
				if matrix[y][x+1] == '*' {
					bombCount++
				}
			}
			// next line
			if !isLastLine {
				if !isFirstCell {
					if matrix[y+1][x-1] == '*' {
						bombCount++
					}
				}
				if matrix[y+1][x] == '*' {
					bombCount++
				}
				if !isLastCell {
					if matrix[y+1][x+1] == '*' {
						bombCount++
					}
				}
			}
			// then finally update the main board
			if bombCount != 0 {
				(*b)[y+1][x+1] = byte(bombCount + 48)
			}
		}
	}

	return nil
}
