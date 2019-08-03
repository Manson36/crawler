package main

import "fmt"

func BubbleSort(arr *[5]int) {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr )- 1 -i; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				(*arr)[j], (*arr)[j + 1] = (*arr)[j + 1], (*arr)[j]
			}
		}
	}

	fmt.Println("排序后arr", *arr)
}

func main() {
	arr := [5]int{32, 44, 4, 33, 6}
	BubbleSort(&arr)
}
