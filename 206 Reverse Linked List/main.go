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

func reverseList(head *ListNode) *ListNode {
	var reverseL *ListNode
	for {
		if head == nil {
			break
		} else {
			var tempList = &ListNode{}
			// 切斷節點
			tempList.Val = head.Val
			tempList.Next = reverseL

			reverseL = tempList

			if head.Next == nil {
				break
			} else {
				head = head.Next
			}
		}
	}

	return reverseL
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

	printLinkList(reverseList(a))

	var b = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	printLinkList(reverseList(b))

	printLinkList(reverseList(&ListNode{}))
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
