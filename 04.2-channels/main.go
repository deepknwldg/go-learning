package main

import (
	"fmt"
)

func echo(arg string, ch chan string) {
	ch <- arg // Send
}

func main() {
	var ch chan string

	fmt.Println(ch) // <nil>

	// unbuffered channel
	chInit := make(chan string) // так как создан канал без размера буфера
	fmt.Println(chInit)         // memory address

	go echo("Hello World", chInit)

	data := <-chInit  // Receive  // то здесь чтение из канала блокируется, до тех пор пока в канал что-нибудь не запишется
	fmt.Println(data) // Hello World

	// buffered channels
	ch1 := make(chan string, 2)
	go echo("value 1", ch1)
	go echo("value 2", ch1)

	val1 := <-ch1
	val2 := <-ch1
	fmt.Println(val1, val2)

	close(ch1)

	var c = make(chan int, 3)
	c <- 20
	c <- 10
	c <- 0
	close(c)
	// здесь читается больше 3 значений, поэтому последние два будут zero value
	for i := 0; i < 5; i++ {
		v, ok := <-c
		fmt.Printf("closed?: %v val: %d\n", !ok, v)
	}

	// send to nil channel blocks forever
	// var c chan string
	// c <- "Hello world!" // Panic: all goroutines are asleep - deadlock!

	// receive from nil channel blocks forever
	// var c chan string
	// fmt.Println(<-c) // Panic: all goroutines are asleep - deadlock!

	// sent to closed channel -> panic
	// c <- "Hello World!"
	// close(c)
	// c <- "Hellow, Panic" // Panic: send on closed channel

	var c2 = make(chan int, 3)
	c2 <- 20
	c2 <- 10
	c2 <- 0
	close(c2)

	for v := range c2 {
		fmt.Printf("value %d\n", v)
	}
}

func directedEcho(arg string, ch chan<- string) {
	ch <- arg // Send only
}

func directedReceive(ch <-chan string) {
	<-ch
}
