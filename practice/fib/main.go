package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+ b
		return a
	}
}

func main() {
	a := fib()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
