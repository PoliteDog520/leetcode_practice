package main

import "log"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if linkLength := getLinkLength(head); linkLength <= 1 {
		return head
	} else {
		linkLength = linkLength / 2
		for i := 0; i < linkLength; i++ {
			head = head.Next
		}
	}

	return head
}

func getLinkLength(linkList *ListNode) int {
	var length int
	if linkList != nil {
		length = 1
		length += getLinkLength(linkList.Next)
	}
	return length
}

func main() {
	var a = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	printLinkList(middleNode(a))
	var b = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	printLinkList(middleNode(b))
}

func printLinkList(list *ListNode) {
	for {
		if list == nil {
			break
		}
		log.Println(list.Val)
		list = list.Next
	}
}
