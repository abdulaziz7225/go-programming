package main

import (
	"fmt"
	"iter"
)

// Example 1
/* func myIterator(yield func(string) bool) {
	vals:= []string{"first", "second"}

	for _, val := range vals {
		if !yield(val) {
			return
		}
	}
} */

// Example 2
/* func myIterator() func(yield func(string) bool) {
	return func(yield func(string) bool) {
		vals := []string{"first", "second"}
		for _, val := range vals {
			if !yield(val) {
				return
			}
		}
	}
} */

// Example 3
func myIterator() iter.Seq[string] {
	return func(yield func(string) bool) {
		vals := []string{"first", "second"}
		for _, val := range vals {
			if !yield(val) {
				return
			}
		}
	}
}

// Example 4
func mySecondIterator() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		vals := []string{"first", "second"}
		for index, val := range vals {
			if !yield(index, val) {
				return
			}
		}
	}
}

func main() {
	for s := range myIterator() {
		fmt.Println(s)
	}

	for i, s := range mySecondIterator() {
		fmt.Println(i, s)
	}
}
