package main

import (
	linkedlist "datastructures/linked_list"
	"fmt"
)

func main() {
	list := linkedlist.ForwardList{}
	list.Show()

	for index := 0; index < 5; index++ {
		list.Prepend(index)
	}

	list.Show()

	for index := 1; index < 10; index++ {
		list.Insert(index*3, index)
	}

	list.Show()

	list.Insert(300, list.Length()-1)

	list.Show()

	for index := 1; index < 8; index++ {
		list.Insert(1234, index*2)
	}

	list.Show()

	if _, err := list.Delete(list.Length()); err != nil {
		fmt.Printf("%v\n", err)
	}

	if _, err := list.Delete(list.Length() - 1); err != nil {
		fmt.Printf("%v\n", err)
	}

	list.Show()

	if _, err := list.Delete(list.Length() - 2); err != nil {
		fmt.Printf("%v\n", err)
	}

	list.Show()

	if _, err := list.Delete(0); err != nil {
		fmt.Printf("%v\n", err)
	}

	list.Show()
}
