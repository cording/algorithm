//  @author: cord
//  @create: 2019-04-18 22:20
// 面试题6：从尾到头打印链表
// 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。
//借助第三方api实现
//思路: 使用栈
package main

import (
	"fmt"
	. "github.com/emirpasic/gods/lists/singlylinkedlist"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/emirpasic/gods/utils"
)

func PrintListReversingly_Iteratively(list *List) {
	stack := arraystack.New()
	iterator := list.Iterator()
	for iterator.Next() {
		stack.Push(iterator.Value())
	}
	for !stack.Empty() {
		val, _ := stack.Pop()
		fmt.Print(val)
	}
	fmt.Println()
}

func PrintListReversingly_Iteratively2(list *List) {
	iterator := list.Iterator()
	for iterator.Next() {
		defer fmt.Print(iterator.Value())
	}
}

func main() {
	list := New()
	list.Add("a")
	list.Append("b")
	list.Prepend("c")
	list.Sort(utils.StringComparator)
	//_, _ = list.Get(0)
	//_, _ = list.Get(100)
	//_ = list.Contains("a", "b", "c")
	//_ = list.Contains("a", "b", "c", "d")
	PrintListReversingly_Iteratively(list)
	fmt.Println("--------------------------------------------")
	PrintListReversingly_Iteratively2(list)
}
