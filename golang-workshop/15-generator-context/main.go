package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func printOut(ctx context.Context) chan string {
	c := make(chan string)

	go func() {
		defer close(c)
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				c <- fmt.Sprintf("Print %v", i)
			}
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c1 := printOut(ctx)
	for v := range c1 {
		fmt.Println(v)
	}
}
