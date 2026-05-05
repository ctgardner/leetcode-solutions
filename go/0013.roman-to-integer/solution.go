// Created by Callum Gardner at 2026/03/12 16:06
// leetgo: dev
// https://leetcode.com/problems/roman-to-integer/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

var numeralValue = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func romanToInt(s string) (ans int) {
	for i := 0; i < len(s); i++ {
		curr := numeralValue[s[i]]
		if i+1 < len(s) && numeralValue[s[i+1]] > curr {
			ans -= curr
		} else {
			ans += curr
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := romanToInt(s)

	fmt.Println("\noutput:", Serialize(ans))
}
