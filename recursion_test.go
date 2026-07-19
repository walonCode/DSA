package main

import "testing"

func TestSumLoop(t *testing.T){
	total := sumLoop([]int{2,4,6})
	if total != 12 {
		t.Fatalf("expected:%v got:%v\n",12, total)
	}
}

func TestSumRecursion(t *testing.T){
	total := sumRecursion([]int{1})
	if total != 1 {
		t.Fatalf("expected:%v, got:%v\n", 1, total)
	}
}


func TestTotal(t *testing.T){
	len := total([]int{1,2,3})
	if len != 3 {
		t.Fatalf("expected:%v, got:%v", 3, len)
	}
}


func TestMax(t *testing.T){
	maxValue := max([]int{2,3,10})
	if maxValue != 10 {
		t.Errorf("expected:%v, got:%v", 10, maxValue)
	}
}