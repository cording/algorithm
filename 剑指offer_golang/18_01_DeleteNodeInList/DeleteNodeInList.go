//  @author: cord
//  @create: 2019-06-05 20:58
// 面试题18（一）：在O(1)时间删除链表结点
// 题目：给定单向链表的头指针和一个结点指针，定义一个函数在O(1)时间删除该结点。

package main

import "fmt"

var nullNode = &Node{}

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	node1 := &Node{"1", nil}
	node2 := &Node{"2", nil}
	node3 := &Node{"3", nil}
	node4 := &Node{"4", nil}
	node5 := &Node{"5", nil}
	node6 := &Node{"6", nil}
	node1.next, node2.next, node3.next, node4.next, node5.next = node2, node3, node4, node5, node6
	DeleteNode(&node1, node3)
	fmt.Printf("delete middle node: %v\n", node1)
}

func test2() {
	node1 := &Node{"11", nil}
	DeleteNode(&node1, node1)
	fmt.Printf("delete head node with total one node: %v\n", node1)
}

func test3() {
	node1 := &Node{"1", nil}
	node2 := &Node{"2", nil}
	node3 := &Node{"3", nil}
	node4 := &Node{"4", nil}
	node5 := &Node{"5", nil}
	node6 := &Node{"6", nil}
	node1.next, node2.next, node3.next, node4.next, node5.next = node2, node3, node4, node5, node6
	DeleteNode(&node1, node6)
	fmt.Printf("delete last node: %v\n", node1)
}

type Node struct {
	value string
	next  *Node
}

//这里要用指针的指针，因为golang是值传递，如果传递节点的指针，然后将其指定为一个空结构体的指针，
// 那么原本的节点不会发生变化，所以只能使用指针的指针，这样修改的是原有结构体指针的值，将其指向空结构体
func DeleteNode(pHead **Node, pToBeDeleted *Node) {
	if pHead == nil || pToBeDeleted == nil {
		return
	}
	//要删除的节点不是尾节点
	//则把下一个节点覆盖当前节点的内容，再把下一个节点置空，这就相当于把当前需要删除的节点删除了
	if pToBeDeleted.next != nil {
		pNext := pToBeDeleted.next
		pToBeDeleted.value = pNext.value
		pToBeDeleted.next = pNext.next
		pNext = nullNode
	} else if *pHead == pToBeDeleted {
		//链表只有一个节点，删除头节点(也是尾节点)
		pToBeDeleted = nullNode
		*pHead = nullNode
	} else {
		//链表有多个节点，删除尾节点
		//由于尾节点的下一个节点为空，所以没法用下一个节点的内容来覆盖当前节点
		//而如果光把尾节点置为空是不够的，还要将上一个节点的next指针指向空，因为它变成尾节点了
		pNode := *pHead
		for pNode.next != pToBeDeleted {
			pNode = pNode.next
		}
		pNode.next = nullNode
		pToBeDeleted = nullNode
	}

}
