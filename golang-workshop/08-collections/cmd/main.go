package main

import "fmt"

func main() {
	array()
	fmt.Println()
	slice()
	fmt.Println()
	maps()
}

func array() {
	strArray := [10]string{"one", "two", "three", "four"}
	fmt.Println(len(strArray), cap(strArray))

	for i, s := range strArray {
		if s != "" {
			fmt.Printf("string %s at position %d\n", s, i)
		}
	}

	strArray[3] = "four2"
	fmt.Println(strArray[3])
}

func slice() {
	strSlice1 := make([]string, 10)
	fmt.Println(len(strSlice1), cap(strSlice1))

	strSlice2 := []string{}
	fmt.Println(len(strSlice2), cap(strSlice2))

	var strSlice3 []string
	fmt.Println(len(strSlice3), cap(strSlice3))

	fmt.Println()

	strSlice1[0] = "one"
	strSlice1[1] = "two"
	strSlice1[2] = "three"
	for i, s := range strSlice1 {
		fmt.Printf("string %s at position %d\n", s, i)
	}

	fmt.Println()

	strSlice2 = append(strSlice2, "one")
	strSlice2 = append(strSlice2, "two", "three")
	strSlice2 = append(strSlice2, "four", "five")
	for i, s := range strSlice2 {
		fmt.Printf("string %s at position %d\n", s, i)
	}
	fmt.Println(len(strSlice2), cap(strSlice2))

	fmt.Println()

	strArray := [6]string{"one", "two", "three", "four", "five", "six"}
	strSlice4 := strArray[2:4]
	fmt.Println(strSlice4, len(strSlice4), cap(strSlice4))
	fmt.Println(strSlice4[0])
	fmt.Println(strSlice4[1])

	strArray[2] = "nothing"
	for i, s := range strSlice4 {
		fmt.Printf("string %s at position %d\n", s, i)
	}
}

func maps() {
	// myMap := map[int]string{}
	var myMap = make(map[int]string)
	myMap[1] = "one"
	myMap[2] = "two"
	myMap[3] = "three"
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	v, ok := myMap[100]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Not found")
	}

	delete(myMap, 100)
	clear(myMap)
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}