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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return sumLink(l1, l2, 0)
}

func sumLink(l1, l2 *ListNode, carry int) *ListNode {
	// 此串列值
	var l1Val, l2Val int
	// 下個鏈結串列
	var l1Next, l2Next *ListNode
	// 計算串列加總
	var sumL = &ListNode{}
	if l1 == nil && l2 == nil {
		if carry > 0 {
			sumL.Val = carry
			return sumL
		} else {
			return nil
		}
	} else if l1 == nil {
		l2Val = l2.Val
		l2Next = l2.Next
	} else if l2 == nil {
		l1Val = l1.Val
		l1Next = l1.Next
	} else {
		l1Val = l1.Val
		l1Next = l1.Next
		l2Val = l2.Val
		l2Next = l2.Next
	}
	sumL.Val = l1Val + l2Val + carry
	// 進位
	if sumL.Val > 9 {
		carry = 1
		sumL.Val -= 10
	} else {
		carry = 0
	}
	sumL.Next = sumLink(l1Next, l2Next, carry)
	return sumL
}

func main() {
	var a = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	var b = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	c := addTwoNumbers(a, b)
	for {
		log.Println(c)
		c = c.Next
		if c == nil {
			break
		}
	}
	//log.Println(addTwoNumbers(a, b))
}
