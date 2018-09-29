/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.
 * @Author: yujinghui
 * @Date: 2018-09-28 14:24:28
 * @Last Modified by:   yujinghui
 * @Last Modified time: 2018-09-28 14:24:28
*/
package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if s == t {
		return true
	}
	if len(s) == len(t) {
		return true
	}
	return false
}

func main242() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("car", "rat"))
	fmt.Println(isAnagram("anadram", "nagaram"))
	fmt.Println(isAnagram("a", "n"))
}
