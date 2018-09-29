/*
Count the number of prime numbers less than a non-negative number, n.

Example:

Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

 * @Author: yujinghui
 * @Date: 2018-09-29 13:12:17
 * @Last Modified by: yujinghui
 * @Last Modified time: 2018-09-29 13:18:00
*/
package main

import (
	"fmt"
)

func countPrimes(n int) int {
	for i := 0; i <= n; i++ {
		fmt.Println(i)
	}
}

func main() {
	countPrimes(4)
}
