//  @author: cord
//  @create: 2019-04-29 23:04
// 面试题8：二叉树的下一个结点
// 题目：给定一棵二叉树和其中的一个结点，如何找出中序遍历顺序的下一个结点？
// 树中的结点除了有两个分别指向左右子结点的指针以外，还有一个指向父结点的指针。

package main

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func GetNext(node *Node) *Node {
	if node == nil {
		return nil
	}
	var next *Node = nil
	if node.right != nil {
		right := node.right
		for right.left != nil {
			right = right.left
		}
		next = right
	} else if node.parent != nil {
		current := node
		parent := node.parent
		//当前节点为右节点
		for parent != nil && current == parent.right {
			node = parent
			parent = parent.parent
		}
		next = parent
	}
	return next
}
