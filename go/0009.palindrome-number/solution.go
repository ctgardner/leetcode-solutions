// Created by Callum Gardner at 2026/03/12 15:32
// leetgo: dev
// https://leetcode.com/problems/palindrome-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPalindrome(x int) bool {
	// negative ints and ints ending with zero, except zero itself, aren't palindromes
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}

	reverse := 0
	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}

	return x == reverse || x == reverse/10
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := isPalindrome(x)

	fmt.Println("\noutput:", Serialize(ans))
}
