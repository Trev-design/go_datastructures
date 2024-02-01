package linkedlist

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
	if index > list.length {
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
