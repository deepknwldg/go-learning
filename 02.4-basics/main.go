package main

import (
	"fmt"
	"time"
)

func main() {
	type Person struct {
		Name, Surname string
		Birthday      time.Time
	}

	date := time.Date(1970, 1, 1, 1, 1, 1, 1, time.Local)

	var p1 Person
	fmt.Println(p1)

	var p2 = Person{
		Name:     "John",
		Surname:  "Wick",
		Birthday: date,
	}
	fmt.Println(p2)

	var p3 = Person{Name: "Somebody"}
	fmt.Println(p3)

	var p4 = Person{"John", "Wick", date}
	fmt.Println(p4.Name)

	ptr := &p4
	fmt.Println(ptr.Name, (*ptr).Name)

	fmt.Println(p2 == *ptr, p2 == p4)

	type SecretPerson struct {
		Name    string
		surname string // не экспортируется в другие пакеты, потому что с маленькой буквы
	}

	type Worker struct {
		Person
		Salary int
	}

	w := Worker{
		Person: Person{
			Name:     "John",
			Surname:  "Wick",
			Birthday: date,
		},
		Salary: 100500,
	}

	fmt.Println(w)
	fmt.Println(w.Name, w.Person.Name)

	type PersonWithTahs struct {
		Name string `json:"name"`
	}

	type Point struct {
		X, Y float64
	}

	point1 := Point{1, 2}
	point2 := point1

	point2.X = 2

	fmt.Println(point1) // {1 2}
	fmt.Println(point2) // {2 2}
}
