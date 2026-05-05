// Created by Callum Gardner at 2026/03/11 16:42
// leetgo: dev
// https://leetcode.com/problems/add-two-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
	}

	return head.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	l1 := Deserialize[*ListNode](ReadLine(stdin))
	l2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := addTwoNumbers(l1, l2)

	fmt.Println("\noutput:", Serialize(ans))
}
