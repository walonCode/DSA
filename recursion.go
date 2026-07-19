package main

import "fmt"

// no base case
func countDown(v int) {
	fmt.Println(v)
	countDown(v - 1)
}

// with base case
func countDownFix(v int) {
	fmt.Println(v)
	if v <= 1 {
		return
	}
	countDownFix(v - 1)
}

// famous factorial
type number interface {
	int | float64
}

func fact[T number](v T) T {
	if v <= 1 {
		return v
	}

	return v * fact(v-1)
}

func loopFact(v int) int {
	if v <= 1 {
		return v
	}

	total := 1

	for i := 1; i <= v; i++ {
		// fmt.Printf("total:%v, value:%v\n", total, i)
		total *= i
	}

	return total
}


//fibonacci
func fibo(v int)int {
	if v <= 1 {
		return v
	}

	return fibo(v-1) + fibo(v-2)
}


func loopFibo(v int) int {
	if v <= 1 {
		return v
	}

	prev, curr := 0, 1

	for i:=2; i <= v; i++ {
		next := prev + curr
 		prev = curr
   		curr = next
	}
	
	
	return curr
}

//finding the sum using a loop
func sumLoop(arr []int)int {
	if len(arr) == 1 {
		return arr[0]
	}
	
	total := 0
	for _, val := range arr {
		total += val
	}

	return total
}

//finding the sum using recursion
func sumRecursion(arr []int)int {	
	if len(arr) == 1 {
		return arr[0]
	}

	first := arr[0]

	return first + sumRecursion(arr[1:])
}

//number items in a list using recursion
func total(arr []int)int {
	if len(arr) == 1 {
		return 1
	}

	len := 1
	return len + total(arr[1:])
}

//maximun number in list using recursion
func max(arr []int)int {
	
	if len(arr) == 1{
		return arr[0]
	}
	
	rest := max(arr[1:])
	if arr[0] > rest {
		return arr[0]
	}
	
	return rest
}
