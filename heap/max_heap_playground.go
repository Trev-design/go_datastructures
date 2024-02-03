package heap

import "fmt"

func ShowUtils() {
	fmt.Println("initialize heap")
	fmt.Println("===============")
	maxHeap := &MaxHeap{}
	fmt.Println(maxHeap)
	fmt.Println()

	fmt.Println("inserting keys in max heap and heapify")
	fmt.Println("======================================")
	for index := 0; index <= 25; index++ {
		maxHeap.Insert(index)
	}
	fmt.Println(maxHeap)
	fmt.Println()

	fmt.Println("extract key and heapify")
	fmt.Println("=======================")
	maxHeap.Extract()
	fmt.Println(maxHeap)
	fmt.Println()
}
