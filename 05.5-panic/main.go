package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var i = rand.Int() % 10
	var a [8]int

	fmt.Println(a[i])

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	defer handlePanic()
	unrecoverableError()

}

func unrecoverableError() {
	panic("stop")
}

func handlePanic() {
	err := recover()
	if err != nil {
		fmt.Printf("recovered from error %s\n", err)
	}
}
