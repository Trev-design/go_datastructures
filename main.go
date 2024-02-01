package main

import (
	linkedlist "datastructures/linked_list"
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
}
