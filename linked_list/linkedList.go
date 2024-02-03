package linked_list

import "fmt"

type node struct {
	value int
	next  *node
}

type ForwardList struct {
	head   *node
	length int
}

func (list *ForwardList) Length() int {
	return list.length
}

func (list *ForwardList) Prepend(value int) {
	newNode := &node{value: value}
	nextNode := list.head
	list.head = newNode
	list.head.next = nextNode
	list.length++
}

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
