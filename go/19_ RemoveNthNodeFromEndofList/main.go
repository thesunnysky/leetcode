package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	before, target, cur := head, head, head
	i := 1
	for cur != nil {
		if i < n {
			target = nil
		} else if i == n {
			target = head
		} else if i == n+1 {
			target = target.Next
			before = head
		} else {
			target = target.Next
			before = before.Next
		}
		cur = cur.Next
		i++
	}
	if n == i-1 {
		head = target.Next
	} else {
		before.Next = target.Next
	}
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	target, cur := head, head
	i := 1
	target = nil
	for cur != nil {
		if i < n {
			target = head
		} else if i == n {
			target = head
		} else {
			target = target.Next
		}
		cur = cur.Next
		i++
	}
	target = target.Next
	return head
}

func main() {
	head := &ListNode{Val: 0, Next: nil}
	rtn := removeNthFromEnd(head, 1)
	fmt.Println(rtn)
}
