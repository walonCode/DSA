package main

func quickSort(arr []int)[]int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := make([]int, 0)
	greater := make([]int, 0)

	for _, val := range arr[1:] {
		if val > pivot {
			greater = append(greater, val)
		}else {
			less = append(less, val)
		}
	}

	result := quickSort(less)
	result = append(result, pivot)
	result = append(result, quickSort(greater)...)
	return result
}