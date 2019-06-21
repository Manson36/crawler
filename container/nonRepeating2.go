package main

import "fmt"

func lengthOfNonRepeatingSunStr2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxlength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch];ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxlength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSunStr2("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSunStr2("bbbbbb"))
	fmt.Println(lengthOfNonRepeatingSunStr2("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSunStr2(""))
	fmt.Println(lengthOfNonRepeatingSunStr2("b"))
	fmt.Println(lengthOfNonRepeatingSunStr2("abcdefg"))
	fmt.Println(lengthOfNonRepeatingSunStr2("这里是哈哈"))
	fmt.Println(lengthOfNonRepeatingSunStr2("一二三二一"))
	fmt.Println(lengthOfNonRepeatingSunStr2(
		"黑化肥挥发会发灰会花飞灰化肥挥发发黑会飞花"))
}
