package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int

	for i :=0; i < 10; i++ {
		go func(i int) {//如果匿名函数不设定参数，那便构成闭包，i计算到10传入，a[10]不存在
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
