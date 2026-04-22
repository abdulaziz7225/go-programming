package main

import (
	"fmt"
)

func main() {
	ip, err := resolveDNS("google.com")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("google.com resolved to: %s\n", ip)
}
