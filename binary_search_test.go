package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T){
	result, steps := binarySearch([]int{1,2,3,4,5}, 4)
	fmt.Println(result, steps)
	if result != 3 {
		t.Fatalf("expected:%v, got:%v", 4, result)
	}

	if steps != 2 {
		t.Fatalf("expected:%v, got:%v", 2, steps)
	}
}