//  @author: cord
//  @create: 2019-05-23 21:45
// 面试题15：二进制中1的个数
// 题目：请实现一个函数，输入一个整数，输出该数二进制表示中1的个数。例如
// 把9表示成二进制是1001，有2位是1。因此如果输入9，该函数输出2。

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NumberOf1_Solution1(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count += 1
		}
		n = n >> 1
	}
	return count
}
func NumberOf1_Solution2(n int) int {
	str := strconv.FormatInt(int64(n), 2)
	str = strings.Replace(str, "0", "", -1)
	return len(str)
}

//上面的解法是O(n)，而还有一种解法是有多少个1循环多少次
//把一个整数减去1和原来的数做与运算，会把该整数最右边的1变为0，
//那么一个整数的二进制中有多少个1就可以进行多少次这样的运算
func NumberOf1_Solution3(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n & (n - 1)
	}
	return count
}

func main() {
	//
	fmt.Println("solution1")
	fmt.Println(NumberOf1_Solution1(9))
	fmt.Println(NumberOf1_Solution1(0))
	fmt.Println(NumberOf1_Solution1(1))
	//
	fmt.Println("solution2")
	fmt.Println(NumberOf1_Solution2(9))
	fmt.Println(NumberOf1_Solution2(0))
	fmt.Println(NumberOf1_Solution2(1))
	//
	fmt.Println("solution3")
	fmt.Println(NumberOf1_Solution3(9))
	fmt.Println(NumberOf1_Solution3(0))
	fmt.Println(NumberOf1_Solution3(1))
}
