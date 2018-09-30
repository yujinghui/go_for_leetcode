/*
The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
6.     312211
7.     13112221
8.     1113213211
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

Example 1:

Input: 1
Output: "1"
Example 2:

Input: 4
Output: "1211"


 * @Author: yujinghui
 * @Date: 2018-09-30 11:47:02
 * @Last Modified by: yujinghui
 * @Last Modified time: 2018-09-30 14:28:53
*/
package main

import (
	"strings"
	"fmt"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	currentStr = countAndSay(n - 1)
	finalStr := ""
	for i, val := range  correntStr {
		if i == strings.Count(s string, substr string)currentStr
	}
}

func main() {
	fmt.Println(countAndSay(4))
}
