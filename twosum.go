// Write by Kotch
// Given an array of integers nums and an integer sum, return the indexes of 2 elements in the array that add up to the given sum.

package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15, 8, 9, 5, 6, 9}
	sum := 18

	// Example 1: Output: 0,1
	// nums := []int{2, 7, 11, 15}
	// sum := 9

	// Example 2: Output: 1,2
	// nums := []int{3, 2, 4}
	// sum := 6

	// Example 3: Output: 0,3
	// nums := []int{3, 1, 2, 3}
	// sum := 6

	// Example 4: Output: no output
	// nums := []int{3, 2}
	// sum := 3

	fmt.Println(sum2(nums, sum))

}

func sum2(nums []int, sum int) []int {

	for i := 0; i <= len(nums)-1; i++ {

		left := i + 1
		right := len(nums) - 1
		sumTmp := sum - nums[i]
		for left <= right {
			mid := left + (right - left)
			if nums[mid] == sumTmp {
				return []int{i, mid}
			} else if nums[mid] < sumTmp {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return []int{}
}
