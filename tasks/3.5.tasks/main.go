package main

import (
	"fmt"
	"go-learning/tasks/3.5.tasks/task1"
	"go-learning/tasks/3.5.tasks/task4"
	"go-learning/tasks/3.5.tasks/task5"
)

func main() {
	fmt.Println(task1.FrequentWord("hi hi he he hi ha hc hb hb"))

	m1 := make(map[int]struct{})
	m1[1] = struct{}{}
	m1[2] = struct{}{}
	m1[3] = struct{}{}
	m1[5] = struct{}{}

	m2 := make(map[int]struct{})
	m2[1] = struct{}{}
	m2[4] = struct{}{}
	m2[3] = struct{}{}

	fmt.Println(task4.MapKeyIntersect(m1, m2))

	slc := [...]string{"Penn", "Teller", "Penn", "Tom", "Teller", "Sam"}

	fmt.Println(task5.ToFrequencyMap(slc[:]))
}
