package matrices

import "fmt"

func MatrixUtils() {
	fmt.Println("Show some diagonal matrix stuff")
	fmt.Println("===============================")

	dima := CreateDiagonalMatrix(7)
	dima.ShowMatrix()
	fmt.Println()

	for index := 0; index < dima.GetDimension(); index++ {
		dima.SetValue(index, index, index+1)
	}

	dima.ShowMatrix()

	fmt.Println()
	fmt.Println("some errors could be return in get and set function")
	fmt.Println()

	if err := dima.SetValue(4, 5, 6); err != nil {
		fmt.Printf("%v\n", err)
	}

	if _, err := dima.GetValue(8, 2); err != nil {
		fmt.Printf("%v\n", err)
	}

	if _, err := dima.GetValue(2, 8); err != nil {
		fmt.Printf("%v\n", err)
	}

	lowma := CreateLowTriMatrix(7)

	for row := 1; row <= 7; row++ {
		for column := 1; column <= 7; column++ {
			if err := lowma.SetValue(row, column, row+column); err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}

	lowma.ShowMatrix()
}
