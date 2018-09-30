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
9.     31131211131221
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

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	saystr := countAndSay(n - 1)
	newStr := ""
	count := 1
	for index := range saystr {
		if index == len(saystr)-2 {
			if saystr[index] != saystr[index+1] {
				newStr = newStr + strconv.Itoa(count) + string(saystr[index]) + "1" + string(saystr[index+1])
			} else {
				newStr = newStr + strconv.Itoa(count+1) + string(saystr[index])
			}
			break
		}
		if saystr[index] != saystr[index+1] {
			newStr = newStr + strconv.Itoa(count) + string(saystr[index])
			count = 1
		} else {
			count++
		}
	}
	return newStr
}

func main() {
	println("--------->" + countAndSay(7))
	// str := "hello, world"
	// println(len(str))
	// for index := range str {
	// print(str[index])
	// }
}
