package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var in = make(chan string)
	var out1 = make(chan string)
	go func() {
		defer close(in)
		in <- "hello. world    . im. bob..   .bob."
	}()

	step1(in, out1)

	var out2 = make(chan string)
	step2(out1, out2)

	var out3 = step3(out2)
	for v := range out3 {
		fmt.Println(v)
	}
}

func step1(in <-chan string, out chan<- string) {
	go func(chanIn <-chan string) {
		defer close(out)

		for value := range chanIn {
			words := strings.Fields(value)
			out <- strings.Join(words, " ")
		}
	}(in)
}

func step2(in <-chan string, out chan<- string) {
	go func(chanIn <-chan string) {
		defer close(out)

		for value := range chanIn {
			sentences := strings.Split(value, ".")

			for _, sentence := range sentences {
				if len(sentence) > 0 {
					out <- strings.Trim(sentence, " ")
				}
			}
		}
	}(in)
}

func step3(in <-chan string) <-chan string {
	out := make(chan string)

	go func(ch <-chan string) {
		defer close(out)

		for value := range ch {
			runes := []rune(value)
			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
				out <- string(runes)
			}
		}
	}(in)

	return out
}
