package matrices

type SquareMatrix interface {
	SetValue(row, column, item int) error
	GetValue(row, column int) (int, error)
	ShowMatrix()
}
