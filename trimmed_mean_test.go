package main

// Import the necessary packages and functions
import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	// Create a new file to write the results
	outputFile, err := os.Create("test_results_Golang.txt")
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	// Define the test cases
	tests := []struct {
		name      string
		input     []interface{}
		lowerTrim float64
		upperTrim float64
		expected  float64
	}{
		{"Test case 1: Floats, Symmetric Trimming", []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}, 0.2, 0.2, 3},
		{"Test case 2: Integers, Lower End Trimming", []interface{}{1, 2, 3, 4, 5}, 0.2, 0, 3.5},
		{"Test case 3: Floats, Upper End Trimming", []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}, 0, 0.2, 2.5},
		{"Test case 4: Floats, No Trimming", []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}, 0, 0, 3},
		{"Test case 5: Integers, Empty Slice", []interface{}{}, 0.1, 0.1, 0},
		{"Test case 6: Test Case With 100 Integers", []interface{}{
			// Repeat the digit -1000 five times. These digits should be trimmed
			-1000, -1000, -1000, -1000, -1000,

			// Type the numbers 1 to 90. This test should calculate the mean of these 90 untrimmed numbers
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
			41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
			51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
			61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
			71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
			81, 82, 83, 84, 85, 86, 87, 88, 89, 90,

			// Repeat the digit 1000000 five times. These digits should be trimmed in this test
			1000000, 1000000, 1000000, 1000000, 1000000}, 0.05, 0.05, 45.5},
		{"Test case 7: Test Case With 100 Floats", []interface{}{
			// Repeat the float -1000.54321 five times. These floats should be trimmed in our test
			-1000.54321, -1000.54321, -1000.54321, -1000.54321, -1000.54321,

			// Type the floats from 0.1 to 9.0 (with step sizes of 0.1). This test should calculate the mean of these 90 untrimmed floats
			0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0,
			1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0,
			2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9, 3.0,
			3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 4.0,
			4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 4.9, 5.0,
			5.1, 5.2, 5.3, 5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 6.0,
			6.1, 6.2, 6.3, 6.4, 6.5, 6.6, 6.7, 6.8, 6.9, 7.0,
			7.1, 7.2, 7.3, 7.4, 7.5, 7.6, 7.7, 7.8, 7.9, 8.0,
			8.1, 8.2, 8.3, 8.4, 8.5, 8.6, 8.7, 8.8, 8.9, 9.0,

			// Repeat the float 3141592.65358979 five times. These floats should be trimmed in our test
			3141592.65358979, 3141592.65358979, 3141592.65358979, 3141592.65358979, 3141592.65358979}, 0.05, 0.05, 4.55},
	}

	// Loop through each test case
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var numbers []float64
			for _, v := range test.input {
				switch vv := v.(type) {
				case float64:
					numbers = append(numbers, vv)
				case int:
					numbers = append(numbers, float64(vv))
				default:
					t.Fatalf("Unexpected type %v in input slice", reflect.TypeOf(v))
				}
			}

			var result float64
			result = trimmedMean(numbers, test.lowerTrim, test.upperTrim)
			//			if test.upperTrim != 0 {
			// Asymmetric trimming
			//				result = trimmedMean(numbers, test.lowerTrim, test.upperTrim)
			//			} else {
			// Symmetric trimming
			//				result = trimmedMean(numbers, test.lowerTrim, test.lowerTrim)
			//			}

			// Write the test result to the output file
			fmt.Fprintf(outputFile, "Test case: %s\n", test.name)
			fmt.Fprintf(outputFile, "Input: %v\n", test.input)
			fmt.Fprintf(outputFile, "Lower Trim: %.2f\n", test.lowerTrim)
			fmt.Fprintf(outputFile, "Upper Trim: %.2f\n", test.upperTrim)
			fmt.Fprintf(outputFile, "Expected: %.2f\n", test.expected)
			fmt.Fprintf(outputFile, "Result: %.2f\n\n", result)

			if result != test.expected {
				t.Errorf("Expected trimmed mean of %v to be %.2f, but got %.2f", test.input, test.expected, result)
			}
		})
	}
}
