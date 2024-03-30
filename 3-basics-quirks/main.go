package main

import "fmt"

func main() {
	ranges()
}

// ranges
func ranges() {
	valuesStr := []string{"a", "b", "c"}
	for index, value := range valuesStr {
		fmt.Println(index, value)
	}

	valuesInt := []int{4, 8, 15, 16, 23, 42}
	for value := range valuesInt {
		fmt.Println(value)
	}
}
