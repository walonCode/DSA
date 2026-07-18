package main

import "fmt"

func main(){

	//testing binary_search function
	list := []int{1,3,5,7,9}
	fmt.Println(binary_search(list, 10))

	//testing selection sort
	{
		list := []int{40, 20, 50, 10, 5, 3 ,1 }
		fmt.Printf("sorted list: %v\n", selectionSort(list))
	}

	//testing linked list 
	{
		list := newLinkedList[int]()
		list.pushFront(1)
		list.pushBack(10)
		list.pushFront(20)

		fmt.Println("linked list: ",list.string())
		fmt.Println("size of the list: ", list.size)
	}
}