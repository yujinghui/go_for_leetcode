/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.


 * @Author: yujinghui
 * @Date: 2018-10-08 09:25:59
 * @Last Modified by: yujinghui
 * @Last Modified time: 2018-10-08 09:30:56
*/
package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	for idx := 0; ; idx++ {
		for i := 0; i < len(strs)-1; i++ {
			if idx > len(strs[i]) || idx > len(strs[i+1]) {
				return strs[0][0 : idx-1]
			}
			if strs[i][0:idx] != strs[i+1][0:idx] {
				return strs[0][0 : idx-1]
			}
		}
	}
}

func main14() {
	strArray := []string{"a"}
	fmt.Println(longestCommonPrefix(strArray))
}
