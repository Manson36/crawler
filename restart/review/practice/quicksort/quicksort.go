package quicksort

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for  j >= p && values[j] >= temp {
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
	if p - left > 1{
		quickSort(values, 0, p)
	}

	if right - p >1 {
		quickSort(values, p + 1, len(values) - 1)
	}
}

func quicksort(values []int) {
	if len(values) <= 1 {
		return
	}
	quickSort(values, 0, len(values)-1)
}

