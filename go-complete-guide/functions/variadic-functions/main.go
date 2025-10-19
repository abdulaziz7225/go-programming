package main

import "fmt"

func main() {
	var sum int = sumUp(1, 10, 15, 63)
	var numbers = []int{34, 10, 54}
	var sum2 int = sumUp(numbers...)
	fmt.Println(sum)
	fmt.Println(sum2)
}

func sumUp(numbers ...int) int {
	var sum int
	for _, value := range numbers {
		sum += value
	}
	return sum
}