package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func processArray(numbers []int) []float64 {
	result := make([]float64, len(numbers))
	for i, num := range numbers {
		result[i] = math.Sqrt(float64(num))
	}
	return result
}

func processArrayParallel(numbers []int, numGoroutines int) []float64 {
	result := make([]float64, len(numbers))
	var wg sync.WaitGroup
	sliceSize := len(numbers) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		start := i * sliceSize
		end := start + sliceSize
		if i == numGoroutines-1 {
			end = len(numbers)
		}
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				result[j] = math.Sqrt(float64(numbers[j]))
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

	start := time.Now()
	_ = processArray(consecutiveNumbers)
	durationWithoutGoroutines := time.Since(start)
	fmt.Println("Time without goroutines:", durationWithoutGoroutines)

	numGoroutines := []int{1, 10, 100, 1000}
	for _, num := range numGoroutines {
		fmt.Printf("Processing with %d goroutines:\n", num)
		start := time.Now()
		_ = processArrayParallel(consecutiveNumbers, num)
		durationWithGoroutines := time.Since(start)
		fmt.Println("Time with", num, "goroutines:", durationWithGoroutines)
	}
}
