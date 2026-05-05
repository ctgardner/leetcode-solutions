// Created by Callum Gardner at 2026/03/13 15:55
// leetgo: dev
// https://leetcode.com/problems/merge-two-sorted-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}

	return head.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	list1 := Deserialize[*ListNode](ReadLine(stdin))
	list2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := mergeTwoLists(list1, list2)

	fmt.Println("\noutput:", Serialize(ans))
}
