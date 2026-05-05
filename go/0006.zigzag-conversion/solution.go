// Created by Callum Gardner at 2026/03/16 14:20
// leetgo: dev
// https://leetcode.com/problems/zigzag-conversion/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func convert(s string, numRows int) string {
	ans := ""
	for row := 0; row < numRows; row++ {
		factor := 1
		j := 0
		for {
			index := j * (numRows-1) + row * factor
			if index >= len(s) {
				break
			}
			ans += string(s[index])

			if factor == -1 {
				factor = 1
			} else {
				j += 2
			}
		}
	}
	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	numRows := Deserialize[int](ReadLine(stdin))
	ans := convert(s, numRows)

	fmt.Println("\noutput:", Serialize(ans))
}

// P   A   H   N
// A P L S I I G
// Y   I   R

// 0   4   8     12
// 1 3 5 7 9  11 13
// 2   6   10

// P     I    N
// A   L S  I G
// Y A   H R
// P     I

// 0     6      12
// 1   5 7    11 13
// 2 4   8 10
// 3     9

// P       H
// A     S I
// Y   I   R
// P L     I G
// A       N

// 0       8
// 1     7 9
// 2   6   10
// 3 5     11 13
// 4       12

// 0N+0          2N+0
// 0N+1     2N-1 2N+1
// 0N+2   2N-2   2N+2
// 0N+3 2N-3     2N+3 4N-3
// 0N+4          2N+4

// row + numRows + (numRows - 2) = gap
// 2 * numRows - 1