// Created by Callum Gardner at 2026/09/08 15:07
// leetgo: dev
// https://leetcode.com/problems/reverse-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseList(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	tail := head.Next
	reversed := reverseList(tail)
	tail.Next = head
	head.Next = nil
	return reversed
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reverseList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
