package main

import "fmt"

func main() {
	//Age := 20 //unused variable error
	
	var _ = "Amy"

	products := map[string]float64{
		"Baby clothing": 284,
		"Handbags":      34,
	}

	//the map returns a string value and a float value, but we are only interested in one value
	for product, _ := range products {
		fmt.Println(product)
	}
}
