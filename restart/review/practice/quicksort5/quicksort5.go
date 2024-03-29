package main

import "fmt"

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}

	quickSort(values, 0, len(values) - 1)
}

func QuickSort2(values []int) {
	if len(values) <= 1 {
		return
	}

	mid, i := values[0], 1
	head, tail := 0, len(values) - 1

	for head < tail {
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
				values[i], values[head] = values[head], values[i]
				head++
				i++
		}
	}

	QuickSort2(values[:head])
	QuickSort2(values[head+1:])
}

func main() {
	var values = []int{22,3,44,5,6,11,8,66}
	QuickSort(values)
	//QuickSort2(values)
	fmt.Println(values)
}