//  @author: cord
//  @create: 2019-05-08 21:52
// 面试题10：斐波那契数列
// 题目：写一个函数，输入n，求斐波那契（Fibonacci）数列的第n项。

package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		fmt.Print(Fibonacci_Solution1(i))
		fmt.Print(" ")
	}
	fmt.Println()
	for i := 0; i < 9; i++ {
		fmt.Print(Fibonacci_Solution2(i))
		fmt.Print(" ")
	}
	fmt.Println()
	for i := 0; i < 9; i++ {
		fmt.Print(Fibonacci_Solution3(0, 1, i))
		fmt.Print(" ")
	}
}

//递归
func Fibonacci_Solution1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci_Solution1(n-1) + Fibonacci_Solution1(n-2)
}

//循环
func Fibonacci_Solution2(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	pre1, pre2 := 0, 1
	for i := 2; i <= n; i++ {
		pre1, pre2 = pre2, pre1+pre2
	}
	return pre2
}

//尾递归
func Fibonacci_Solution3(pre1 int, pre2 int, n int) int {
	if n <= 0 {
		return pre1
	}
	if n == 1 {
		return pre2
	}
	n--
	return Fibonacci_Solution3(pre2, pre1+pre2, n)
}
