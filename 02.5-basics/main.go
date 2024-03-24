package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name     string
	Birthday time.Time
}

// func (variable T) methodName(params) (returnTypes) {}

func (c Person) doNothing() {

}

func (c Person) CanDriveCar() bool {
	return c.Birthday.AddDate(18, 0, 0).Before(time.Now())
}

func (c Person) UpdateName(name string) {
	c.Name = name
}

func (c *Person) UpdateNameByPointer(name string) {
	c.Name = name
}

type MyInt int

func (i MyInt) isGreater(value int) bool {
	return i > MyInt(value)
}

func main() {
	const layout = "2006-Jan-02"
	tm, _ := time.Parse(layout, "2000-Sep-03")
	p := Person{
		Name:     "John",
		Birthday: tm,
	}
	fmt.Println(p)

	p.UpdateName("new name")
	fmt.Println(p) // unchanged

	p.UpdateNameByPointer("new name") // go сам понимает, что нужно сделать разыменование (&p).UpdateNameByPointer(...)
	fmt.Println(p)                    // changed
	fmt.Println(p.CanDriveCar())      // true

	i := MyInt(10)
	fmt.Println(i.isGreater(10)) // false
}
