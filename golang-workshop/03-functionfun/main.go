package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	fmt.Println(secondFunc(firstFunc()))
	fmt.Println(secondFunc(thirdFunc()()))
	fmt.Println(fourthFunc(thirdFunc())())
}

func firstFunc() (int, int) {
	return 2, 3
}

func secondFunc(a, b int) int {
	return a + b
}

func thirdFunc() func() (int, int) {
	return firstFunc
}

func fourthFunc(a func() (int, int)) func() int {
	b, c := a()

	return func() int {
		return b + c
	}
}
