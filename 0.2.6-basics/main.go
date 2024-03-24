package main

import (
	"fmt"
)

func main() {
	var arr [4]int
	fmt.Println(arr[0])
	fmt.Println(arr[3])
	// fmt.Println(arr[4]) invalid index

	fmt.Println(arr)

	var arr2 = [4]int{10, 20, 30, 40}
	fmt.Println(arr2)

	arr3 := [4]int{10, 20, 30, 40}
	fmt.Println(arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Printf("Index: %d, Element: %d\n", i, arr3[i])
	}

	for i, value := range arr3 {
		fmt.Printf("Index: %d, Element: %d\n", i, value)
	}

	for _, value := range arr3 {
		fmt.Printf("Element: %d\n", value)
	}

	for i := range arr3 {
		fmt.Printf("Index: %d\n", i)
	}

	for range arr3 {
		fmt.Println("?")
	}

	multiDimArr := [...][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, el := range multiDimArr { // где value равен одномерному массиву
		fmt.Printf("Index: %d, Element: %d\n", i, el)
	}

	a1 := [3]string{"One", "Two", "Three"}
	var b1 = a1 // copy of a1

	b1[0] = "One!!!"
	fmt.Println(a1)
	fmt.Println(b1)
}
