//  @author: cord
//  @create: 2019-05-06 21:37
// 面试题9：用两个栈实现队列
// 题目：用两个栈实现一个队列。队列的声明如下，请实现它的两个函数appendTail
// 和deleteHead，分别完成在队列尾部插入结点和在队列头部删除结点的功能。
//思路：删除元素的步骤，如果stack2不为空，stack2栈顶的元素是最先进入队列的元素，可以弹出;
//当stack2为空的时候，将stack1中的元素逐渐弹出并压入stack2中，并弹出stack2栈顶元素
//新增元素的话就直接往stack1里面加就行了，这时就算stack2不为空，弹出操作也是优先弹出stack2里的元素,
//所以直接压入stack1即可

package main

import (
	"fmt"
	. "github.com/emirpasic/gods/stacks/arraystack"
)

type Queue struct {
	stack1 *Stack
	stack2 *Stack
}

func main() {
	queue := NewQueue()
	queue.deleteHead()
	for i := 0; i < 10; i++ {
		queue.appendTail(i)
	}
	queue.deleteHead()
	queue.deleteHead()
	queue.printQueue()
}

func NewQueue() (queue *Queue) {
	return &Queue{New(), New()}
}

//尾部添加
func (queue *Queue) appendTail(element interface{}) {
	queue.stack1.Push(element)
}

//删除头部节点
func (queue *Queue) deleteHead() (head interface{}) {
	if queue.stack2.Size() == 0 {
		for queue.stack1.Size() != 0 {
			p, _ := queue.stack1.Pop()
			queue.stack2.Push(p)
		}
	}
	if queue.stack2.Size() == 0 {
		fmt.Println("queue is empty")
	}
	head, _ = queue.stack2.Pop()
	return head
}

//打印队列,会把队列消费光
func (queue *Queue) printQueue() {
	stack1 := queue.stack1
	stack2 := queue.stack2
	for i := 0; i < stack1.Size(); i++ {
		p, _ := stack1.Pop()
		fmt.Print(p)
		fmt.Print(" ")
	}
	for stack2.Size() != 0 {
		p, _ := stack2.Pop()
		fmt.Print(p)
		fmt.Print(" ")
	}
	fmt.Println()

}
