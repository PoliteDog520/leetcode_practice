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

func detectCycle(head *ListNode) *ListNode {

	var (
		linkAddress = make(map[*ListNode]struct{})
	)
	return checkCycleLinkList(linkAddress, head)
}

func checkCycleLinkList(linkAddr map[*ListNode]struct{}, linkList *ListNode) *ListNode {
	// 檢查linkList為空或下個linklist為空，則確認非環狀連結串列
	if linkList == nil || linkList.Next == nil {
		return nil
	} else {
		// 判斷此連結串列是否已出現過，若出現則為環狀連結串列的起始點
		if _, ok := linkAddr[linkList]; ok {
			return linkList
		} else {
			linkAddr[linkList] = struct{}{}
			return checkCycleLinkList(linkAddr, linkList.Next)
		}
	}
}

func main() {
	//var head = ListNode{
	//	Val: 3,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 0,
	//			Next: &ListNode{
	//				Val:  4,
	//				Next: nil,
	//			},
	//		},
	//	},
	//}
	//
	//head.Next.Next.Next.Next = head.Next

	log.Println(detectCycle(nil))
}
