package main

import "fmt"

func main() {
	m := "asdfsdfsa"
	for i, v := range []byte(m) {
		fmt.Printf("i=%d, v= %X\n", i, v)
	}

	n := "Yes, 今天好开心！"
	for i, v := range []byte(n) {
		fmt.Printf("i=%d, v= %X\n", i, v)
	}
}

