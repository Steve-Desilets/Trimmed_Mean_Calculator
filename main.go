package main

import (
	"fmt"
	"sort"
)

func trimmedMean(numbers []float64, lowerTrim float64, upperTrim ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	// Sort the numbers in ascending order
	sort.Float64s(numbers)

	// Calculate the number of elements to trim from each end
	lowerTrimCount := int(lowerTrim * float64(len(numbers)))
	var upperTrimCount int
	if len(upperTrim) > 0 {
		upperTrimCount = int(upperTrim[0] * float64(len(numbers)))
	} else {
		upperTrimCount = lowerTrimCount
	}

	// Calculate the sum of the elements after trimming
	sum := 0.0
	for i := lowerTrimCount; i < len(numbers)-upperTrimCount; i++ {
		sum += numbers[i]
	}

	// Calculate and return the trimmed mean
	return sum / float64(len(numbers)-lowerTrimCount-upperTrimCount)
}

func main() {
	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lowerTrim := 0.2
	upperTrim := 0.4 // Optional

	fmt.Printf("Trimmed mean with symmetric trimming: %.2f\n", trimmedMean(numbers, lowerTrim))
	fmt.Printf("Trimmed mean with asymmetric trimming: %.2f\n", trimmedMean(numbers, lowerTrim, upperTrim))
}
