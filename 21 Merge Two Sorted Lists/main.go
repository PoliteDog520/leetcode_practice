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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var mergeList = &ListNode{}
	var linkHead = mergeList
	for {
		log.Println(list1.Val, list2.Val)
		if list1.Val <= list2.Val {
			mergeList.Val = list1.Val
			if list1.Next == nil {
				mergeList.Next = list2
				break
			} else {
				list1 = list1.Next
			}
		} else {
			mergeList.Val = list2.Val
			if list2.Next == nil {
				mergeList.Next = list1
				break
			} else {
				list2 = list2.Next
			}
		}
		mergeList.Next = &ListNode{}
		mergeList = mergeList.Next
	}

	return linkHead
}

func main() {
	var l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	var l2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	log.Println(mergeTwoLists(l1, l2))
}
