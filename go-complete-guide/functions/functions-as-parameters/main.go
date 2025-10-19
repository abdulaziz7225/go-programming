package main

import "fmt"

type transformFn func(int) int

func main() {
	var numbers = []int{1, 2, 3, 4}
	// var doubled = transformNumbers(&numbers, double)
	// var tripled = transformNumbers(&numbers, func(num int) int {
	// 	return num * 3
	// })

	var double = createTransformer(2)
	var triple = createTransformer(3)
	var doubled = transformNumbers(&numbers, double)
	var tripled = transformNumbers(&numbers, triple)

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(nums *[]int, transform transformFn) []int {
	var result []int

	for _, value := range *nums {
		result = append(result, transform(value))
	}

	return result
}

// func double(num int) int {
// 	return num * 2
// }

// func triple(num int) int {
// 	return num * 2
// }

func createTransformer(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}