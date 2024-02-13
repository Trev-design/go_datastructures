package main

import (
	"datastructures/heap"
	"datastructures/linked_list"
	"datastructures/matrices"
)

func main() {
	linked_list.ShowUtils()
	heap.ShowUtils()
	diMa := matrices.Create(7)
	diMa.ShowMatrix()
}
