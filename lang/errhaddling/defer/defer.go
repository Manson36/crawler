package main

import (
	"bufio"
	"fmt"
	"github.com/improve01/errorhanding/fib"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func WriteFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL | os.O_CREATE, 0666)
	if err != nil{
		if patherr, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(patherr.Op, patherr.Path, patherr.Err)
		}
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	WriteFile("fib.txt")
	//tryDefer()
}