// Created by Callum Gardner at 2026/09/03 15:42
// leetgo: dev
// https://leetcode.com/problems/reverse-vowels-of-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseVowels(s string) string {
	vowels := []byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	vowelsLookup := make([]bool, 256)
	for _, i := range vowels {
		vowelsLookup[i] = true
	}

	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; {
		for i < j && !vowelsLookup[b[i]] {
			i++
		}
		for j > i && !vowelsLookup[b[j]] {
			j--
		}
		if i != j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}
	return string(b)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := reverseVowels(s)

	fmt.Println("\noutput:", Serialize(ans))
}
