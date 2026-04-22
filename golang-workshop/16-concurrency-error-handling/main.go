package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func task(ctx context.Context, id int) error {
	timer := time.NewTimer(time.Duration(rand.Intn(1000)) * time.Millisecond)
	defer timer.Stop()

	select {
	case <-timer.C:
		if id%3 == 0 {
			return fmt.Errorf("task %d failed", id)
		}
		fmt.Printf("Task %d completed\n", id)
		return nil
	case <-ctx.Done():
		return context.Cause(ctx)
	}
}

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	for i := 1; i <= 10; i++ {
		id := i
		g.Go(func() error {
			return task(ctx, id)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Group failed with error: %v\n", err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}
