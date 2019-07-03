package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastoccurred := make(map[rune]int)
	start := 0
	maxlength := 0

	for i, ch := range []rune(s) {
		lastI, ok := lastoccurred[ch]
		if ok && lastI > start {
			start = lastI + 1
		}

		if i - start + 1  > maxlength {
			maxlength = i-start+1
		}
		lastoccurred[ch] = i
	}
	return maxlength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(
			"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
