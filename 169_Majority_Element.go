/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2

 * @Author: yujinghui
 * @Date: 2018-09-27 20:44:36
 * @Last Modified by: yujinghui
 * @Last Modified time: 2018-09-28 14:30:55
*/
package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	countmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		count,_ := countmap[key]
		countmap[key] = 1 + count
	}

	for k, v := range countmap {
		if v > len(nums)/2 {
			return k
		}
	}
	return 1
}

func main() {
	nums := []int{1}
	fmt.Println(majorityElement(nums))
}
