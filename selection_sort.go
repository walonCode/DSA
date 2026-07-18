package main

import "cmp"

func findSmallest[T cmp.Ordered](arr []T) int {
	smallestIdx := 0
	for i, val := range arr {
		if val < arr[smallestIdx] {
			smallestIdx = i
		}
	}
	return smallestIdx
}

func selectionSort[T cmp.Ordered](arr []T) []T {
	newArr := make([]T, 0, len(arr))
	copiedArr := make([]T, len(arr))
	copy(copiedArr, arr)

	for len(copiedArr) > 0 {
		idx := findSmallest(copiedArr)
		newArr = append(newArr, copiedArr[idx])
		copiedArr = append(copiedArr[:idx], copiedArr[idx+1:]...)
	}
	return newArr
}