/*

Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?


 * @Author: yujinghui
 * @Date: 2018-09-29 09:29:39
 * @Last Modified by: yujinghui
 * @Last Modified time: 2018-09-29 13:15:17
*/
package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	var missingNumber = len(nums)
	for index := range nums {
		missingNumber = missingNumber ^ index ^ nums[index]
	}
	return missingNumber
}

func main268() {
	fmt.Println(missingNumber([]int{0}))
}
