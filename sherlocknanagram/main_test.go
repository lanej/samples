package main

import (
	"reflect"
	"testing"
)

// https://www.hackerrank.com/challenges/minimum-swaps-2/problem

// Function Description

// Complete the function minimumSwaps in the editor below. It must return an integer representing the minimum number of swaps to sort the array.

// minimumSwaps has the following parameter(s):

// arr: an unordered array of integers
// Input Format

// The first line contains an integer, , the size of .
// The second line contains  space-separated integers .

// Constraints

// Output Format

// Return the minimum number of swaps to sort the given array.

// Sample Input 0

// 4
// 4 3 1 2
// Sample Output 0

// 3
// Explanation 0

// Given array
// After swapping  we get
// After swapping  we get
// After swapping  we get
// So, we need a minimum of  swaps to sort the array in ascending order.

// Sample Input 1

// 5
// 2 3 4 1 5
// Sample Output 1

// 3
// Explanation 1

// Given array
// After swapping  we get
// After swapping  we get
// After swapping  we get
// So, we need a minimum of  swaps to sort the array in ascending order.

// Sample Input 2

// 7
// 1 3 5 2 4 6 8
// Sample Output 2

// 3
// Explanation 2

// Given array
// After swapping  we get
// After swapping  we get
// After swapping  we get
// So, we need a minimum of  swaps to sort the array in ascending order.

var cases = []struct {
	input  string
	output int32
}{
	{
		input:  "abba",
		output: int32(4),
	},
	{
		input:  "abc",
		output: int32(0),
	},
	{
		input:  "ifailuhkqq",
		output: int32(3),
	},
	{
		input:  "kkkk",
		output: int32(10),
	},
}

func TestAnagrams(t *testing.T) {
	for x, i := range cases {
		ans := sherlockAndAnagrams(i.input)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
