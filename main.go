package main

import (
	"fmt"
)

type Integer interface {
	int | int8 | int16 | int32 | int64
}

type Float interface {
	float32 | float64
}

// app function
func add[T Integer | Float](a, b T) T {
	return a + b
}

// main function that will be called when the application starts
func main() {
	fmt.Println("Starting print1...")
	fmt.Println(add(1, 2))
	fmt.Println(add(3.34, 1.64))
}
