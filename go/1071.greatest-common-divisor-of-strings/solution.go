// Created by Callum Gardner at 2026/09/03 08:43
// leetgo: dev
// https://leetcode.com/problems/greatest-common-divisor-of-strings/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func gcdOfStrings(str1 string, str2 string) string {
	g := gcd(len(str1), len(str2))
	str3 := str1[:g]

	s1 := strings.Repeat(str3, len(str1)/g)
	s2 := strings.Repeat(str3, len(str2)/g)
	if s1 == str1 && s2 == str2 {
		return str3
	}

	return ""
}

// gcd returns the greatest common denominator of integers a and b, calculated using Euclid's Division Algorithm.
func gcd(a, b int) int {
	for a != 0 && b != 0 {
		a, b = b, a%b
	}
	if a == 0 {
		return b
	} else {
		return a
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	str1 := Deserialize[string](ReadLine(stdin))
	str2 := Deserialize[string](ReadLine(stdin))
	ans := gcdOfStrings(str1, str2)

	fmt.Println("\noutput:", Serialize(ans))
}
