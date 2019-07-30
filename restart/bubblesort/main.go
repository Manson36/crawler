package main

import "fmt"

func test() {
	s := make([]int, 5)
	fmt.Printf("cap =%v\n", cap(s))

	s = append(s, 10)
	s = append(s, 11)
	x := append(s, 12, 13)
	y := s
	y = append(y, 14)
	fmt.Println(s, x, y)
	fmt.Printf("cap = %v", cap(s))
}

func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr", *arr)
	temp := 0

	for i := 0; i < len(arr) -1 ; i++{
		for j := 0; j < len(arr) - 1- i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	}

	fmt.Println("排序后arr", *arr)
}

func main() {
	//test()
	arr := [5]int{23,44,6,22,2}
	BubbleSort(&arr)
}
