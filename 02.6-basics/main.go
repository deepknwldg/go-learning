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

	// slices
	var zeroSlice []int
	fmt.Println(zeroSlice, zeroSlice == nil, len(zeroSlice)) // [] true 0

	var oneElementSlice = make([]int, 1, 10)
	fmt.Println(oneElementSlice, len(oneElementSlice), cap(oneElementSlice)) // [0] 1 10

	sliceWithLenEqCap := make([]int, 5)
	fmt.Println(sliceWithLenEqCap, len(sliceWithLenEqCap), cap(sliceWithLenEqCap)) // [0 0 0 0 0] 5 5

	myFavSlice := []string{"I", "like", "learning", "Go"}
	fmt.Println(myFavSlice)

	fmt.Println(myFavSlice[1:2]) // like
	fmt.Println(myFavSlice[:3])  // I like learning

	myArr := [5]int{20, 15, 5, 30, 25}
	mySlice := myArr[1:4]

	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", myArr, len(myArr), cap(myArr))
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", mySlice, len(mySlice), cap(mySlice))

	for _, v := range myFavSlice {
		fmt.Println(v)
	}

	s1 := []string{"a", "b", "c", "d"}
	s2 := make([]string, len(s1))

	e := copy(s2, s1)

	fmt.Println("Src:", s1)
	fmt.Println("Dst:", s2)
	fmt.Println("Elements:", e)

	s3 := make([]string, 2)
	e2 := copy(s3, s1)

	fmt.Println(s3, e2)

	s3 = append(s3, "c")
	fmt.Println(s3)

	s3 = append(s3, "d", "e")
	fmt.Println(s3)
}
