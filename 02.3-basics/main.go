package main

import (
	"fmt"
)

func main() {
	var pointer *int

	fmt.Println(pointer) // nil

	value := 43
	pointer = &value
	fmt.Println(pointer)  // 0x1400010e010
	fmt.Println(*pointer) // 42

	*pointer += 10
	fmt.Println(*pointer) // 54
	fmt.Println(value)
}
