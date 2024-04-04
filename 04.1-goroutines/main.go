package main

import (
	"fmt"
	"time"
)

func echo(arg string) {
	fmt.Println(arg)
}

// Если нужно передать данные в рамках процесса то передайте их копию, а не ссылку на оригинальное значение
func main() {
	go echo("Hello World")
	time.Sleep(time.Millisecond)
}
