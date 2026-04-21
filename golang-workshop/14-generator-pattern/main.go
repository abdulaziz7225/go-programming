package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printOut() <-chan string {
	c := make(chan string)

	go func() {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Print %v", i)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	c := printOut()
	for range 10 {
		fmt.Println(<-c)
	}
}
