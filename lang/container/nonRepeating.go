package main

import "fmt"

func lengthOfNonRepeatingSunStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxlength := 0

	for i, ch := range []byte(s) {
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
	fmt.Println(lengthOfNonRepeatingSunStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSunStr("bbbbbb"))
	fmt.Println(lengthOfNonRepeatingSunStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSunStr(""))
	fmt.Println(lengthOfNonRepeatingSunStr("b"))
	fmt.Println(lengthOfNonRepeatingSunStr("abcdefg"))
}
