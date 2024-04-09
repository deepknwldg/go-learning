package main

import "fmt"

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

	data := <-chInit // Receive  // то здесь чтение из канала блокируется, до тех пор пока в канал что-нибудь не запишется

	fmt.Println(data) // Hello World

	// buffered channels
	ch1 := make(chan string, 2)
	go echo("value 1", ch1)
	go echo("value 2", ch1)

	val1 := <-ch1
	val2 := <-ch1
	fmt.Println(val1, val2)

	close(ch1)
}
