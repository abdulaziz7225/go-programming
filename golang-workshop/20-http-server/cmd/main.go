package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", handler)

	log.Println("Server started ...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
}
