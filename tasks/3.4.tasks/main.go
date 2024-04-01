package main

import (
	"fmt"
	"go-learning/tasks/3.4.tasks/task1"
)

func main() {
	fmt.Println(task1.Reverse([]int{1, 2, 3, 4}))
	fmt.Println(task1.RemoveDuplicates([]int{1, 5, 2, 3, 5, 4, 1}))
}
