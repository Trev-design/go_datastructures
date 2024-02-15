package linked_list

import "fmt"

type CompareFunc func(int, int) bool

type node struct {
	value int
	next  *node
}

type ForwardList struct {
	head   *node
	length int
}

// creates a pointer of a new list
func Create() *ForwardList {
	return &ForwardList{}
}

// returns the length of a list
func (list *ForwardList) Length() int {
	return list.length
}

// insert a new value at the beginning of a list
func (list *ForwardList) Prepend(value int) {
	newNode := &node{value: value}
	nextNode := list.head
	list.head = newNode
	list.head.next = nextNode
	list.length++
}

// insert a new value an any place of a list
func (list *ForwardList) Insert(value int, index int) error {
	if index >= list.length {
		return fmt.Errorf("index %d to big", index)
	}

	if index < 0 {
		return fmt.Errorf("invalid index")
	}

	newNode := &node{value: value}

	if index == 0 {
		newNode.next = list.head
		list.head = newNode
		list.length++
	} else {
		current := list.head

		for place := 0; place < index; place++ {
			current = current.next
		}

		newNode.next = current.next
		current.next = newNode
		list.length++
	}

	return nil
}

// deletes the value at given index of a list
func (list *ForwardList) Delete(index int) (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("list is empty")
	}

	if index >= list.length {
		return 0, fmt.Errorf("index %d is to big", index)
	}

	if index < 0 {
		return 0, fmt.Errorf("index must be positive")
	}

	var value int

	if index == 0 {
		value = list.head.value
		list.head = list.head.next
		list.length--
		return value, nil
	}

	current := list.head

	for place := 0; place < index-1; place++ {
		current = current.next
	}

	value = current.next.value
	current.next = current.next.next
	list.length--
	return value, nil
}

// reverses a list
func (list *ForwardList) Reverse() error {
	if list.head == nil {
		return fmt.Errorf("ther is nothing to reverse")
	}
	current := list.head
	var shift, reverseNode *node = nil, nil

	for current != nil {
		reverseNode = shift
		shift = current
		current = current.next
		shift.next = reverseNode
	}
	list.head = shift

	return nil
}

func (list *ForwardList) Sort(compare CompareFunc) {
	mergeSort(&list.head, compare)
}

func mergeSort(source **node, compare CompareFunc) {
	head := *source
	if head == nil || head.next == nil {
		return
	}

	var left, right *node

	split(*source, &left, &right)

	mergeSort(&left, compare)
	mergeSort(&right, compare)

	*source = sortedMerge(left, right, compare)
}

func split(source *node, left, right **node) {
	slow := source
	fast := source.next

	for fast != nil {
		fast = fast.next

		if fast != nil {
			slow = slow.next
			fast = fast.next
		}
	}

	*left = source
	*right = slow.next
	slow.next = nil
}

func sortedMerge(left, right *node, compare CompareFunc) *node {
	var result *node = nil

	if left == nil {
		return right
	} else if right == nil {
		return left
	}

	if compare(left.value, right.value) {
		result = left
		result.next = sortedMerge(left.next, right, compare)
	} else {
		result = right
		result.next = sortedMerge(left, right.next, compare)
	}

	return result
}

// checks if the list is sorted with a specific pattern
func (list *ForwardList) IsSorted(compare CompareFunc) bool {
	current := list.head

	for current.next != nil {
		if !compare(current.value, current.next.value) {
			return false
		}

		current = current.next
	}
	return true
}

// append a list to another
func (list *ForwardList) Concat(other *ForwardList) error {
	if list.head == nil {
		return fmt.Errorf("you can not concatenate a list on an empty list")
	}

	current := list.head
	otherHead := other.head

	for current.next != nil {
		current = current.next
	}

	current.next = otherHead
	otherHead = nil

	list.length += other.length

	return nil
}

// merges to sorted lists together
func (list *ForwardList) Merge(other *ForwardList, compare CompareFunc) error {
	current, otherCurrent := list.head, other.head

	if !list.IsSorted(compare) {
		return fmt.Errorf("the list must be sorted to make a merge")
	}

	if !other.IsSorted(compare) {
		return fmt.Errorf("the list which you want to use to make a merge is not sorted")
	}

	var newNode, lastNode *node

	if compare(current.value, otherCurrent.value) {
		lastNode = current
		newNode = lastNode
		current = current.next
		newNode.next = nil
	} else {
		lastNode = otherCurrent
		newNode = lastNode
		otherCurrent = otherCurrent.next
		newNode.next = nil
	}

	for current != nil && otherCurrent != nil {
		if compare(current.value, otherCurrent.value) {
			lastNode.next = current
			lastNode = current
			current = current.next
			lastNode.next = nil
		} else {
			lastNode.next = otherCurrent
			lastNode = otherCurrent
			otherCurrent = otherCurrent.next
			lastNode.next = nil
		}
	}

	if current != nil {
		lastNode.next = current
	}

	if otherCurrent != nil {
		lastNode.next = otherCurrent
	}

	list.length += other.length

	return nil
}

// calculates the sum of a list
func (list *ForwardList) Sum() int {
	current := list.head
	sum := 0
	for current != nil {
		sum += current.value
		current = current.next
	}

	return sum
}

// calculates the average of the list
func (list *ForwardList) Average() float64 {
	return float64(list.Sum()) / float64(list.length)
}

// prints a list well formated
func (list *ForwardList) Show() {
	toPrint := list.head
	fmt.Print("[ ")
	for toPrint != nil {
		if toPrint.next != nil {
			fmt.Printf("%d, ", toPrint.value)
		} else {
			fmt.Printf("%d ", toPrint.value)
		}
		toPrint = toPrint.next
	}
	fmt.Println("]")
}
