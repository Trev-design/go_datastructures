package linked_list

import (
	"fmt"
)

func ShowUtils() {
	fmt.Println("Initialize a list")
	fmt.Println("=================")
	list := Create()
	list.Show()

	fmt.Println()

	fmt.Println("Prepend items to a list")
	fmt.Println("=======================")
	for index := 0; index < 5; index++ {
		list.Prepend(index)
	}

	list.Show()

	fmt.Println()

	fmt.Println("Insert items on various places")
	fmt.Println("==============================")
	for index := 1; index < 10; index++ {
		list.Insert(index*3, index)
	}

	list.Show()
	fmt.Println()

	fmt.Println("Insert an item at the end of the list")
	fmt.Println("=====================================")
	list.Insert(300, list.Length()-1)

	list.Show()
	fmt.Println()

	fmt.Println("insert items on varius places")
	fmt.Println("=============================")
	for index := 1; index < 8; index++ {
		list.Insert(1234, index*2)
	}

	list.Show()
	fmt.Println()

	fmt.Println("delete an item failed")
	fmt.Println("=====================")
	if _, err := list.Delete(list.Length()); err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println()

	fmt.Println("delete an item at the end succeed")
	fmt.Println("=================================")
	if _, err := list.Delete(list.Length() - 1); err != nil {
		fmt.Printf("%v\n", err)
	}

	list.Show()
	fmt.Println()

	fmt.Println("delete an item from everywhere")
	fmt.Println("==============================")
	if _, err := list.Delete(list.Length() - 2); err != nil {
		fmt.Printf("%v\n", err)
	}

	list.Show()
	fmt.Println()

	fmt.Println("delete an item from beginning")
	fmt.Println("=============================")
	if _, err := list.Delete(0); err != nil {
		fmt.Printf("%v\n", err)
	}

	list.Show()
	fmt.Println()

	fmt.Println("reversing a linked list")
	fmt.Println("=======================")
	list.Reverse()

	list.Show()
	fmt.Println()

	list2 := Create()

	for index := 0; index < 5; index++ {
		list2.Prepend(8)
	}

	fmt.Println("concatenating two linked lists")
	fmt.Println("==============================")
	list.Concat(list2)

	list.Show()
	fmt.Println()

	list.Sort(func(a int, b int) bool { return a <= b })
	list.Show()
	fmt.Println()

	list.Sort(func(a int, b int) bool { return a >= b })
	list.Show()
	fmt.Println()

	fmt.Println("merging two linked lists")
	fmt.Println("========================")

	list3 := Create()
	list4 := Create()
	list5 := Create()
	list6 := Create()
	list7 := Create()
	list8 := Create()
	list9 := Create()

	for index := 5; index > 0; index-- {
		list3.Prepend(index)
		list4.Prepend(index + 1)
		list5.Prepend(index + 2)
		list6.Prepend(index * 2)
		list7.Prepend(index * 3)
		list8.Prepend(index + 4)
	}

	for index := 15; index > 0; index-- {
		list9.Prepend(index)
	}

	list3.Show()
	list4.Show()
	list5.Show()
	list6.Show()
	list7.Show()
	list8.Show()
	list9.Show()

	list3.Merge(list4, func(a int, b int) bool { return a <= b })
	list3.Show()

	list3.Merge(list5, func(a int, b int) bool { return a <= b })
	list3.Show()

	list3.Merge(list6, func(a int, b int) bool { return a <= b })
	list3.Show()

	list3.Merge(list7, func(a int, b int) bool { return a <= b })
	list3.Show()

	list3.Merge(list8, func(a int, b int) bool { return a <= b })
	list3.Show()

	list3.Merge(list9, func(a int, b int) bool { return a <= b })
	list3.Show()

	fmt.Printf("%d\n", list3.length)

	fmt.Printf("The average is: %f\n", list3.Average())
}
