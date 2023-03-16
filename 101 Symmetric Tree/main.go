package main

import "log"

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var leftQueue = make(chan *TreeNode, 1000)
	var rightQueue = make(chan *TreeNode, 1000)

	leftTreeToQueue(leftQueue, root.Left)
	rightTreeToQueue(rightQueue, root.Right)
	if len(leftQueue) != len(rightQueue) {
		return false
	}
	subTreeHigh := len(leftQueue)
	for i := 0; i < subTreeHigh; i++ {
		leftElem, _ := <-leftQueue
		rightElem, _ := <-rightQueue
		if leftElem == nil && rightElem == nil {
			continue
		} else if leftElem == nil {
			return false
		} else if rightElem == nil {
			return false
		} else {
			if leftElem.Val != rightElem.Val {
				return false
			}
		}
	}

	return true
}

// leftTreeToQueue Convert Tree to Queue
func leftTreeToQueue(queue chan *TreeNode, leaf *TreeNode) {
	if leaf == nil {
		queue <- nil
		return
	}
	if leaf.Left == nil && leaf.Right == nil {
		queue <- leaf
	} else {
		leftTreeToQueue(queue, leaf.Left)
		queue <- leaf
		rightTreeToQueue(queue, leaf.Right)
	}
}

// rightTreeToQueue Convert Tree to Queue
func rightTreeToQueue(queue chan *TreeNode, leaf *TreeNode) {
	if leaf == nil {
		queue <- nil
		return
	}
	if leaf.Left == nil && leaf.Right == nil {
		queue <- leaf
	} else {
		rightTreeToQueue(queue, leaf.Right)
		queue <- leaf
		leftTreeToQueue(queue, leaf.Left)
	}
}

func main() {
	var testTree = &TreeNode{}

	testTree.Val = 1
	//testTree.Left = &TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val: 3,
	//	},
	//	Right: &TreeNode{
	//		Val: 4,
	//	},
	//}
	//testTree.Right = &TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val: 4,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//	},
	//}

	testTree.Left = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
	}
	testTree.Right = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
	}

	log.Println(isSymmetric(testTree))
}
