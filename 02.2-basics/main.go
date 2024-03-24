package main

import (
	"fmt"
	"go-learning/02.2-basics/example"
	alias "go-learning/02.2-basics/example2"
)

func main() {
	fmt.Println(example.MyExportedValue)
	fmt.Println(alias.MyOtherValue)
}
