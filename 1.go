package main

import "fmt"

func main() {
	var consecutiveNumbers [1000]int

	for i := 0; i < 1000; i++ {
		consecutiveNumbers[i] = i + 1
	}

	fmt.Println(consecutiveNumbers)
}
