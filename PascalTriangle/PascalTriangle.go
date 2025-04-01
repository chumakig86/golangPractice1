package main

import "fmt"

// Function to calculate binomial coefficient without recursion
func binomialCoeff(n, k int) int {
	if k > n-k {
		k = n - k
	}
	res := 1
	for i := 0; i < k; i++ {
		res *= (n - i)
		res /= (i + 1)
	}
	return res
}

// Function to print Pascal's Triangle
func printPascalsTriangle(n int) {
	for i := 0; i < n; i++ {
		// Add spaces for centering the row
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		// Print elements of the row
		for j := 0; j <= i; j++ {
			fmt.Print(binomialCoeff(i, j), " ")
		}
		// Move to the next line
		fmt.Println()
	}
}

func main() {
	var n int
	// Read input from the user
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&n)

	// Generate and print Pascal's Triangle
	printPascalsTriangle(n)
}
