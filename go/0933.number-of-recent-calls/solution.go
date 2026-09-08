// Created by Callum Gardner at 2026/09/08 12:34
// leetgo: dev
// https://leetcode.com/problems/number-of-recent-calls/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

const threshold = 3000

type RecentCounter struct {
	timestamps []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (r *RecentCounter) Ping(t int) (ans int) {
	minTimestamp := t - threshold
	minIndex := -1
	i, j := 0, len(r.timestamps)
	for i < j {
		k := i + (j-i)/2
		if minTimestamp == r.timestamps[k] {
			minIndex = k
			break
		}

		if minTimestamp > r.timestamps[k] {
			i = k + 1
		} else {
			j = k
		}
	}
	if minIndex == -1 {
		minIndex = j
	}
	r.timestamps = append(r.timestamps[minIndex:], t)
	return len(r.timestamps)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "ping":
			methodParams := MustSplitArray(params[i])
			t := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Ping(t))
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
