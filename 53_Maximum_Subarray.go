/*

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.


--------------------------------------------------------------------------------------------------------------------------------------

solution1:

solution2:


 * @Author: yujinghui
 * @Date: 2018-09-30 10:24:18
 * @Last Modified by: yujinghui
 * @Last Modified time: 2018-09-30 11:43:54
*/
package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	var largestSum float64 = -math.MaxFloat64
	var currentSum float64 = -math.MaxFloat64
	for i := 0; i < len(nums); i++ {
		currentSum = math.Max(currentSum+float64(nums[i]), float64(nums[i]))
		largestSum = math.Max(largestSum, currentSum)
	}
	return int(largestSum)
}

func main() {
	fmt.Println(maxSubArray([]int{-1}))
}
