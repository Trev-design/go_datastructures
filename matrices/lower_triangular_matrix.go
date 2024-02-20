package matrices

import (
	"fmt"
	"math"
)

type LowTriMatrix struct {
	container []int
	dimension int
}

func CreateLowTriMatrix(dimension int) *LowTriMatrix {
	return &LowTriMatrix{
		container: make([]int, dimension*(dimension+1)/2),
		dimension: dimension,
	}
}

func (matrix *LowTriMatrix) SetValue(row, column, item int) error {
	if row > matrix.dimension {
		return fmt.Errorf("row index to big")
	}

	if column > matrix.dimension {
		return fmt.Errorf("column index to big")
	}

	if row >= column {
		matrix.container[row*(row-1)/2+column-1] = item
	}

	return nil
}

func (matrix *LowTriMatrix) GetValue(row, column int) (int, error) {
	if row >= matrix.dimension {
		return math.MinInt, fmt.Errorf("row index to big")
	}

	if column >= matrix.dimension {
		return math.MinInt, fmt.Errorf("column index to big")
	}

	if row >= column {
		return matrix.container[row*(row-1)/2+column-1], nil
	}

	return math.MinInt, nil
}

func (matrix *LowTriMatrix) ShowMatrix() {
	for row := 1; row <= matrix.dimension; row++ {
		for column := 1; column <= matrix.dimension; column++ {
			if row >= column {
				fmt.Printf(" %d", matrix.container[row*(row-1)/2+column-1])
			} else {
				fmt.Print(" 0")
			}
		}
		fmt.Println()
	}
}
