package main

import (
	"fmt"

	"golang.source-fellows.com/training/carman/v3/internal/http"
	"golang.source-fellows.com/training/carman/v3/internal/memory"
)

func main() {
	repository := &memory.CarRepository{}
	fmt.Println(http.StartServer(repository))
}
