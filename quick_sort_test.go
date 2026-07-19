package main

import "testing"

func TestQuickSort(t *testing.T){
	arr := []int{10,40,20,1,5,6,7,30,2,1}
	sortedArr := quickSort(arr)

	if sortedArr[0] != 1 {
		t.Errorf("expected:%v, got:%v", 1, sortedArr[0])
	}
}