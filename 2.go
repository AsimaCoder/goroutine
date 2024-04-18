package main

import (
	"fmt"
	"math"
)

func processNumbers(numbers []int) []float64 {
	result := make([]float64, len(numbers))
	for i, num := range numbers {
		result[i] = math.Sqrt(float64(num))
	}
	return result
}

func main() {

	numbers := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = i + 1
	}

	processedNumbers := processNumbers(numbers)

	fmt.Println(processedNumbers)
}
