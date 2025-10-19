package main

import "fmt"


type Product struct {
	title string
	id int
	price float64
}

func main() {

	var hobbies = [3]string{"Cycling", "Running", "Video games"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	var slice1 = hobbies[0:2]
	var slice2 = hobbies[:2]
	fmt.Println(slice1, slice2)
	fmt.Println(slice1, len(slice1), cap(slice1))
	
	slice1 = slice1[1:3]
	fmt.Println(slice1, len(slice1), cap(slice1))

	// 5)
	courseGoals := []string{"Learn Go", "Build microservice with Go"}
	fmt.Println(courseGoals)

	var products = []Product{
		{
			"laptop",
			1,
			1000.99,
		},
		{
			"standing-desk",
			2,
			419.99,
		},
	}

	fmt.Println(products)
	products = append(products, Product{
		"ergonomic chair",
		3,
		350.0,
	})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
