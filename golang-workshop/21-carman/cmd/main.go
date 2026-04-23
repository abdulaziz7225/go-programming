package main

import (
	"fmt"

	"golang.source-fellows.com/training/carman/internal/http"
	"golang.source-fellows.com/training/carman/internal/memory"
)

func main() {
	repository := &memory.CarRepository{}
	fmt.Println(http.StartServer(repository))
}
