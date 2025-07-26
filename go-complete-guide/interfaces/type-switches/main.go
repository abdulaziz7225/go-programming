package main

import "fmt"

// func printOutput(data any) {
func printOutput(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("Integer:", data)
	case float64:
		fmt.Println("Float:", data)
	case rune:
		fmt.Println("Rune (single 'int32' character):", data)
	case string:
		fmt.Println("String:", data)
	}
}

func printOutput2(data any) {
	intVal, ok := data.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := data.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	runeVal, ok := data.(rune)
	if ok {
		fmt.Println("Rune (single 'int32' character):", runeVal)
		return
	}

	stringVal, ok := data.(string)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}
}

// func printOutputCaller(functionName func(any)) {
func printOutputCaller(functionName func(interface{})) {
	functionName(1)
	functionName(2.75)
	functionName(true)
	functionName('A')
	functionName("Gopher")
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
func main() {
	printOutputCaller(printOutput)
	printOutputCaller(printOutput2)

	result := add(2, 3)
	fmt.Println(result)
}
