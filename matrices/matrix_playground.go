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

	if err := dima.SetValue(4, 5, 6); err != nil {
		fmt.Printf("%v\n", err)
	}

	if _, err := dima.GetValue(8, 2); err != nil {
		fmt.Printf("%v\n", err)
	}

	if _, err := dima.GetValue(2, 8); err != nil {
		fmt.Printf("%v\n", err)
	}
}
