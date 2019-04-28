//  @author: cord
//  @create: 2019-04-18 22:20
// 面试题6：从尾到头打印链表
// 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。
//思路: 栈，defer，递归
package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Node struct {
	value string
	next  *Node
}

func PrintListReversingly_Iteratively(head *Node) {
	stack := arraystack.New()
	for head != nil {
		stack.Push(head.value)
		head = head.next
	}
	for !stack.Empty() {
		val, _ := stack.Pop()
		fmt.Print(val)
	}
	fmt.Println()
}

func PrintListReversingly_Iteratively2(head *Node) {
	defer fmt.Println()
	for head != nil {
		defer fmt.Print(head.value)
		head = head.next
	}
}

func PrintListReversingly_Iteratively3(head *Node) {
	if head != nil {
		PrintListReversingly_Iteratively3(head.next)
		fmt.Print(head.value)
	}
}

func main() {
	head := Node{"a", nil}
	head.next = &Node{"b", nil}
	head.next.next = &Node{"c", nil}
	PrintListReversingly_Iteratively(&head)
	fmt.Println("--------------------------------------------")
	PrintListReversingly_Iteratively2(&head)
	fmt.Println("--------------------------------------------")
	PrintListReversingly_Iteratively3(&head)
}
