package prefix_sum

import (
	"fmt"
	"log"
)

func Launch() {
	input := []int{1, 2, 3, 4, 5, 6}
	i := 1
	j := 3
	log.Printf("prefixSum input: %v, %d, %d \n", input, i, j)
	output := prefixSum(input, i, j)
	fmt.Printf("prefixSum output between positions %d and %d: %v \n", i, j, output)
}

// Prefix Sum involves preprocessing an array to create a new array where each element at index i represents the sum of the array from the start up to i. This allows for efficient sum queries on subarrays.
// Use this pattern when you need to perform multiple sum queries on a subarray or need to calculate cumulative sums.
// Sample Problem:
// Given an array nums, answer multiple queries about the sum of elements within a specific range [i, j].
// Example:
// Input: nums = [1, 2, 3, 4, 5, 6], i = 1, j = 3
// Output: 9
// Explanation:
// Preprocess the array A to create a prefix sum array: P = [1, 3, 6, 10, 15, 21].
// To find the sum between indices i and j, use the formula: P[j] - P[i-1].Prefix Sum involves preprocessing an array to create a new array where each element at index i represents the sum of the array from the start up to i. This allows for efficient sum queries on subarrays.
// Use this pattern when you need to perform multiple sum queries on a subarray or need to calculate cumulative sums.
// Sample Problem:
// Given an array nums, answer multiple queries about the sum of elements within a specific range [i, j].
// Example:
// Input: nums = [1, 2, 3, 4, 5, 6], i = 1, j = 3
// Output: 9
// Explanation:
// Preprocess the array A to create a prefix sum array: P = [1, 3, 6, 10, 15, 21].
// To find the sum between indices i and j, use the formula: P[j] - P[i-1].
// https://blog.algomaster.io/p/15-leetcode-patterns
func prefixSum(input []int, i, j int) int {
	// prefixSum calculates the sum of elements in the input slice between indices i and j (inclusive).
	// It uses a prefix sum array to efficiently compute the sum in constant time after preprocessing.
	//
	// Parameters:
	// - input: The input slice of integers
	// - i: The starting index (1-based) for the sum calculation
	// - j: The ending index (1-based) for the sum calculation
	//
	// Returns:
	// - The sum of elements between indices i and j, or 0 if the input is invalid
	//
	// Time complexity: O(n) for preprocessing, O(1) for each query
	// Space complexity: O(n) for the prefix sum array
	inputLen := len(input)
	if inputLen == 0 {
		log.Println("input has no content!")
		return 0
	}

	if i < 0 || j < 0 {
		log.Println("input has no content!")
		return 0
	}

	prefixSum := make([]int, inputLen)
	for k := range input {
		if k < inputLen {
			for _, b := range input[:k+1] {
				prefixSum[k] += b
			}
		}
	}
	log.Println("prefixSum: ", prefixSum)

	if j <= inputLen {
		if i-1 < 0 {
			return prefixSum[j]
		}

		if i-1 < j {
			return prefixSum[j] - prefixSum[i-1]
		}
	}
	return 0
}
