package main

import (
	"fmt"
	"go-learning/tasks/3.7.tasks/task1"
)

func main() {
	a := 1
	b := 5
	fmt.Println(a, b)
	task1.Swap(&a, &b)
	fmt.Println(a, b)
}
