//  @author: cord
//  @create: 2019-04-24 23:19
// 面试题7：重建二叉树
// 题目：输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输
// 入的前序遍历和中序遍历的结果中都不含重复的数字
//假如前序遍历是 {1,2,4,5,6,3,7,8}，中序遍历是 {4,2,6,5,1,3,8,7}，
// 前序遍历中的第一个值就是根节点，这个节点把中序遍历分成了两部分，
// 分别对应根节点的左右子树，则左子树的中序遍历是左边部分 (4,2,6,5)，
// 一共4个数字，所以左子树的前序遍历是根节点 1 之后的 4 个数字 (2,4,5,6)。
// 于是进入递归分析左子树，左子树分析完之后用同样的方式分析右子树即可
package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func ConstructTree(preorder []int, inorder []int) *Node {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := Node{}
	root.Value = preorder[0]
	var index int
	for i, element := range inorder {
		if element == preorder[0] {
			index = i
			break
		}
	}
	index++ //slice的右边界值是开区间,不包含
	if index >= 1 {
		root.Left = ConstructTree(preorder[1:index], inorder[:index])
		root.Right = ConstructTree(preorder[index:], inorder[index:])
	}
	return &root
}

func main() {
	preorder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inorder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	node := ConstructTree(preorder, inorder)
	printNode(node)

}

func printNode(node *Node) {
	if node != nil {
		fmt.Println(node.Value)
		if node.Left != nil {
			printNode(node.Left)
		}
		if node.Right != nil {
			printNode(node.Right)
		}
	}
}
