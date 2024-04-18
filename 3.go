package main

import (
	"fmt"
	"math"
	"sync"
)

func processArray(num []int, numGoroutines int) []float64 {
	result := make([]float64, len(num))
	var wg sync.WaitGroup
	sliceSize := len(num) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		start := i * sliceSize
		end := start + sliceSize
		if i == numGoroutines-1 {
			end = len(num)
		}
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				result[j] = math.Sqrt(float64(num[j]))
			}
		}(start, end)
	}
	wg.Wait()
	return result
}

func main() {
	consecutiveNumbers := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		consecutiveNumbers[i] = i + 1
	}
	numGoroutines := []int{1, 10, 100, 1000}
	for _, num := range numGoroutines {
		fmt.Printf("Processing with %d goroutines:\n", num)
		processedNumbers := processArray(consecutiveNumbers, num)
		fmt.Println(processedNumbers[:5], "...", processedNumbers[len(processedNumbers)-5:])
		fmt.Println()
	}
}
