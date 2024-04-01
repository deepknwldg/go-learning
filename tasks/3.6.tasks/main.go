package main

import (
	"fmt"
	"go-learning/tasks/3.6.tasks/task1"
	"go-learning/tasks/3.6.tasks/task2"
	"go-learning/tasks/3.6.tasks/task3"
	"go-learning/tasks/3.6.tasks/task4"
)

func main() {
	str := "The quick brown 狐 狐 jumped over the lazy 犬"
	fmt.Println("Before:", str)
	fmt.Println("After:", task1.Reverse(str))

	fmt.Println(task2.RemoveSpaces("text rest   best  post   "))
	fmt.Println(task3.FrequentRune("reere rre 狐狐狐狐狐狐"))

	fmt.Println(task4.StringLengthWithoutSpaces("Test part \t hi "))
}
