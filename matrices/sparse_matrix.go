package matrices

import "fmt"

type element struct {
	row, column, value int
}

type SparseMatrix struct {
	rows     int
	columns  int
	nonZeros int
	elements []*element
}

func CreateSparseMatrix(rows, columns int) *SparseMatrix {
	return &SparseMatrix{rows: rows, columns: columns}
}

func (matrix *SparseMatrix) InsertElement(row, column, value int) error {
	if row > matrix.rows {
		return fmt.Errorf("row to big")
	}

	if column > matrix.columns {
		return fmt.Errorf("column to big")
	}

	matrix.elements = append(matrix.elements, &element{row: row, column: column, value: value})
	matrix.nonZeros++

	return nil
}

func (matrix *SparseMatrix) DisplayMatrix() {
	index := 0

	for row := 0; row < matrix.rows; row++ {
		for column := 0; column < matrix.columns; column++ {
			if row == matrix.elements[index].row && column == matrix.elements[index].column {
				fmt.Printf("%d ", matrix.elements[index].value)
				index++
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Printf("\n")
	}
}
