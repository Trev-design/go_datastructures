package matrices

import "fmt"

type DiagonalMatrix struct {
	container []int
	dimension int
}

func Create(dimension int) *DiagonalMatrix {
	return &DiagonalMatrix{
		container: make([]int, dimension),
		dimension: dimension,
	}
}

func (matrix *DiagonalMatrix) SetValue(row, column, item int) error {
	if row >= matrix.dimension {
		return fmt.Errorf("row index to big")
	}

	if column >= matrix.dimension {
		return fmt.Errorf("column index to big")
	}

	if row != column {
		return fmt.Errorf("your index is not diagonal")
	}

	matrix.container[row] = item
	return nil
}

func (matrix *DiagonalMatrix) GetValue(row, column int) (int, error) {
	if row >= matrix.dimension {
		return -1, fmt.Errorf("row index to big")
	}

	if column >= matrix.dimension {
		return -1, fmt.Errorf("column index to big")
	}

	if row != column {
		return -1, fmt.Errorf("your index is not diagonal")
	}

	return matrix.container[row], nil
}

func (matrix *DiagonalMatrix) ShowMatrix() {
	for row := 0; row < matrix.dimension; row++ {
		for column := 0; column < matrix.dimension; column++ {
			if row == column {
				fmt.Printf(" %d", matrix.container[row])
			} else {
				fmt.Print(" 0")
			}
		}
		fmt.Println()
	}
}
