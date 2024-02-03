package linked_list

import "fmt"

func ShowUtils() {
	fmt.Println("Initialize a list")
	fmt.Println("=================")
	list := ForwardList{}
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
}
