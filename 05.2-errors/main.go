package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) int {
	return a / b
}

func divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}

	return a / b, nil
}

func divide3(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionZero
	}

	return a / b, nil
}

var ErrDivisionZero = errors.New("division by 0")

type RetryError struct {
	NumRetries int
}

func NewRetryError(num int) RetryError {
	return RetryError{NumRetries: num}
}

func (r RetryError) Error() string {
	return fmt.Sprintf("Retries: %d", r.NumRetries)
}

type SomeOtherError struct {
	NumRetries int
}

func (s SomeOtherError) Error() string {
	return fmt.Sprintf("Other error with retries: %d", s.NumRetries)
}

func main() {
	// fmt.Println(divide(10, 0)) // panic: runtime error: integer divide by zero

	val, err := divide2(10, 0)
	if err != nil {
		if err.Error() == "division by 0" {
			fmt.Println("Please provide valid input")
		} else {
			fmt.Printf("Some unknown error: %s\n", err)
		}
	} else {
		fmt.Printf("%d / %d = %d\n", 10, 0, val)
	}

	// sentinel error
	val, err = divide3(10, 0)
	if err != nil {
		if errors.Is(err, ErrDivisionZero) {
			fmt.Println("Please provide valid input")
		} else {
			fmt.Printf("Some unknown error: %s\n", err)
		}
	} else {
		fmt.Printf("%d / %d = %d\n", 10, 0, val)
	}

	// custom errors
	err = NewRetryError(3)
	fmt.Println(err) // Retries: 3

	var retryError RetryError
	if errors.As(err, &retryError) {
		fmt.Println(retryError)
	}

	var otherError SomeOtherError
	if errors.As(err, &otherError) {
		fmt.Println(otherError)
	}

	if e, ok := err.(RetryError); ok {
		fmt.Println("type assert err", e) // type assert err Retries: 3
	}

	if e, ok := err.(SomeOtherError); ok { // ok = false
		fmt.Println("type assert err", e) // wouldn't execute
	}
}
