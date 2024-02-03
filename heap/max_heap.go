package heap

import "fmt"

type MaxHeap struct {
	array []int
}

func (maxHeap *MaxHeap) Insert(key int) {
	maxHeap.array = append(maxHeap.array, key)
	maxHeap.maxHeapifyUp(len(maxHeap.array) - 1)
}

func (maxHeap *MaxHeap) Extract() (int, error) {
	extracted := maxHeap.array[0]

	if len(maxHeap.array) == 0 {
		return -1, fmt.Errorf("heap is empty")
	}

	length := len(maxHeap.array) - 1
	maxHeap.array[0] = maxHeap.array[length]
	maxHeap.array = maxHeap.array[:length]

	maxHeap.maxHeapifyDown(0)

	return extracted, nil
}

func (maxHeap *MaxHeap) maxHeapifyUp(index int) {
	for maxHeap.array[parent(index)] < maxHeap.array[index] {
		maxHeap.swap(parent(index), index)
		index = parent(index)
	}
}

func (maxHeap *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(maxHeap.array) - 1
	left, right := leftIndex(index), rightIndex(index)
	childToCompare := 0

	for left <= lastIndex {
		if left == lastIndex {
			childToCompare = left
		} else if maxHeap.array[left] > maxHeap.array[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}

		if maxHeap.array[index] < maxHeap.array[childToCompare] {
			maxHeap.swap(index, childToCompare)
			index = childToCompare
			left, right = leftIndex(index), rightIndex(index)
		} else {
			return
		}
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftIndex(index int) int {
	return 2*index + 1
}

func rightIndex(index int) int {
	return 2*index + 2
}

func (maxHeap *MaxHeap) swap(index1, index2 int) {
	maxHeap.array[index1], maxHeap.array[index2] = maxHeap.array[index2], maxHeap.array[index1]
}
