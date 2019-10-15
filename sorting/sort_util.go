package sorting

func SelectionSort(values []int) {
	for i := 0; i < len(values); i++ {
		minIndex := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[minIndex] {
				minIndex = j
			}
		}
		values[i], values[minIndex] = values[minIndex], values[i]
	}
}

func InsertionSort(values []int) {
	for i := 0; i < len(values); i++ {
		for j := i; j > 0; j-- {
			if values[j] < values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
			} else {
				break
			}
		}
	}
}

func BubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		swapHappened := false
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				swapHappened = true
			}
		}
		if !swapHappened {
			return
		}
	}
}

func CountSort(values []int, max int) {
	counts := make([]int, max)
	for _, value := range values {
		counts[value]++
	}

	counter := 0
	for value, count := range counts {
		for i := 0; i < count; i++ {
			values[counter] = value
			counter++
		}
	}
}

// TODO: determine why current implementation of MergeSort so slow
func MergeSort(values []int) []int {
	c := make(chan []int)
	go mergeSort(values, c)
	return <-c
}

func mergeSort(values []int, resChan chan []int) {
	if len(values) == 1 {
		resChan <- values
		return
	}

	center := len(values) / 2
	left := values[:center]
	right := values[center:]
	c := make(chan []int, 2)
	go mergeSort(left, c)
	go mergeSort(right, c)
	resChan <- merge(<-c, <-c)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	if leftIndex < len(left) {
		result = append(result, left[leftIndex:]...)
	}
	if rightIndex < len(right) {
		result = append(result, right[rightIndex:]...)
	}
	return result
}

func QuickSort(values []int) []int {
	quickSort(values, 0, len(values)-1)
	return values
}

func quickSort(values []int, l int, r int) {
	if l < r {
		q := partition(values, l, r)
		quickSort(values, l, q)
		quickSort(values, q+1, r)
	}
}

func partition(values []int, l int, r int) int {
	center := (l + r) / 2
	centerValue := values[center]
	i := l
	j := r
	for i <= j {
		for values[i] < centerValue {
			i++
		}
		for values[j] > centerValue {
			j--
		}
		if i >= j {
			break
		}
		values[i], values[j] = values[j], values[i]
		i++
		j--
	}
	return j
}
