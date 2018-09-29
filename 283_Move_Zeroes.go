/*
 * Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

 * Example:

 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 * @Author: yujinghui
 * @Date: 2018-09-27 16:55:38
 * @Last Modified by: yujinghui
 * @Last Modified time: 2018-09-27 20:47:24
 */
package main

import "fmt"
// my solution
func moveZeroes(nums []int) {
	lastNonZeroAt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroAt] = nums[i]
			lastNonZeroAt++
		}
	}
	for i := lastNonZeroAt; i < len(nums); i++ {
		nums[i] = 0
	}
}


func main283() {
	nums := []int{0, 0, 1}
	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}

