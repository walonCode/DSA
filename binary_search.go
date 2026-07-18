package main

import "cmp"

//run at logarithm time big O(logn)
func binary_search[T cmp.Ordered](arr []T, item T) (int, int){
	low := 0
	step := 0 
	high := len(arr) - 1

	for low <= high {
		step++
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == item {
			return mid, step
		}else if guess > item {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	return -1, step
}

