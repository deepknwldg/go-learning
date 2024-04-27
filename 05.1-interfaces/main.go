package main

import (
	"fmt"
	"io"
	"os"
)

type NetworkSender struct {
}

func (s *NetworkSender) Write(p []byte) (n int, err error) {
	return fmt.Printf("Network send: %s", string(n))
}

type SampleInterface interface {
	Method() string
}

type OtherInterface interface {
	SampleInterface
	OtherInterface() string
}

type BidirectionalCommunication struct {
}

func (b *BidirectionalCommunication) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (b *BidirectionalCommunication) Write(p []byte) (n int, err error) {
	return 0, nil
}

type SampleStruct struct {
}

func (s *SampleStruct) Method() string {
	return ""
}

func main() {
	var w io.Writer
	w = os.Stdout
	_, _ = w.Write([]byte("Hello, world\n"))

	w = &NetworkSender{}
	_, _ = w.Write([]byte("Hello, world\n"))

	// differences
	var i interface{} = "hello"
	s := i.(string) // type assertion
	fmt.Println(s)

	// number := i.(int) // panic: interface conversion: interface {} is string, not int
	number, ok := i.(int)
	if !ok {
		fmt.Println("i is not an int")
	} else {
		fmt.Println(number)
	}

	// type conversion
	intVar := string(97)
	fmt.Println(intVar) // a

	var int32Var int32 = 100500
	var int64Var int64

	int64Var = int64(int32Var)
	fmt.Println(int64Var)

	// type switch
	var t interface{} = "hello"
	switch t := t.(type) {
	case string:
		fmt.Printf("string: %s\n", t)
	case bool:
		fmt.Printf("string: %s\n", t)
	case int:
		fmt.Printf("string: %s\n", t)
	default:
		fmt.Printf("I don't know about type %T!\n", t)
	}

	// zero value
	var sample SampleInterface
	fmt.Println(sample, sample == nil) // <nil> true

	var sampleStruct *SampleStruct = nil
	sample = sampleStruct
	fmt.Println(sampleStruct, sampleStruct == nil) // <nil> false

	var x io.ReadWriter = &BidirectionalCommunication{}
	_, _ = x.Read([]byte{})
	_, _ = x.Write([]byte{})

	// var y OtherInterface
	// y.Method() // panic: runtime error: invalid memory address or nil pointer dereference
	// y.OtherInterface()
}
