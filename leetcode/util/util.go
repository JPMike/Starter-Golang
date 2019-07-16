package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetOrderedListNode(count int, dummyHead bool) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	cur := head
	for i := 1; i <= count; i++ {
		node := &ListNode{Val: i, Next: nil}
		cur.Next = node
		cur = node
	}
	if dummyHead {
		return head
	}
	return head.Next
}

func PrintListNode(head *ListNode, dummyHead bool) {
	if dummyHead {
		head = head.Next
	}
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print(" ")
		head = head.Next
	}
	fmt.Print("\n")
}

func PrintNums(s []int) {
	for _, v := range s {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}
