package main

import "testing"

func TestSelectionSort(t *testing.T){
	arr := []int{10, 4, 3, 20, 15, 4, 0, 1}
	sorted_arr := selectionSort(arr)
	if sorted_arr[0] != 0 {
		t.Errorf("expected:%v, got:%v", 0, sorted_arr[0])
	}
}